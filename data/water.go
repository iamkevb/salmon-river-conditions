package data

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type SiteData struct {
	Number    string
	Name      string
	Latitude  float64
	Longitude float64
	Times     []string
	Readings  []int32
	Timestamp time.Time
}

type apiWaterData struct {
	Value apiWaterDataValue
}
type apiWaterDataValue struct {
	TimeSeries []apiWaterTimeseries
}
type apiWaterTimeseries struct {
	SourceInfo apiWaterSourceInfo
	Variable   apiWaterVariable
	Values     []apiWaterValue
}
type apiWaterSourceInfo struct {
	SiteName    string
	GeoLocation apiWaterGeoLocation
}
type apiWaterGeoLocation struct {
	GeogLocation apiWaterGeoLocationChild
}
type apiWaterGeoLocationChild struct {
	Latitude  float64
	Longitude float64
}
type apiWaterVariable struct {
	Unit apiWaterUnit
}
type apiWaterUnit struct {
	Unitcode string
}
type apiWaterValue struct {
	Value []apiWaterValueReading
}
type apiWaterValueReading struct {
	Value    string
	DateTime string
}

var sitesMap = make(map[string]*SiteData)
var waterMutex sync.Mutex

func GetSiteData(usgsCode string) *SiteData {
	waterMutex.Lock()
	defer waterMutex.Unlock()

	data, ok := sitesMap[usgsCode]
	if !ok || data.Timestamp.Add(time.Hour).Before(time.Now()) {
		data = fetchWaterData(usgsCode)
	}
	return data
}

func fetchWaterData(usgsCode string) *SiteData {
	bytes := loadSiteData(usgsCode)
	var apiData apiWaterData
	err := json.Unmarshal(bytes, &apiData)
	if err != nil {
		fmt.Println("Error parsing water data: ", err.Error())
		return nil
	}
	for i, ts := range apiData.Value.TimeSeries {
		if ts.Variable.Unit.Unitcode != "ft3/s" {
			continue
		}

		siteData := &SiteData{
			Number:    usgsCode,
			Name:      apiData.Value.TimeSeries[i].SourceInfo.SiteName,
			Latitude:  apiData.Value.TimeSeries[i].SourceInfo.GeoLocation.GeogLocation.Latitude,
			Longitude: apiData.Value.TimeSeries[i].SourceInfo.GeoLocation.GeogLocation.Longitude,
			Timestamp: time.Now(),
		}
		var times []string
		var readings []int32
		for _, v := range ts.Values[0].Value {
			parsedTime, _ := time.Parse(time.RFC3339, v.DateTime)
			formatted := parsedTime.Format("1/2, 3:04pm")
			times = append(times, formatted)

			parsedReading, _ := strconv.ParseInt(v.Value, 10, 32)
			readings = append(readings, int32(parsedReading))
		}
		siteData.Times = times
		siteData.Readings = readings
		return siteData
	}

	return nil
}

func loadSiteData(usgsCode string) []byte {
	url := fmt.Sprintf("https://waterservices.usgs.gov/nwis/iv/?format=json&sites=%s&period=P5D&siteStatus=all", usgsCode)

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return []byte{}
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return []byte{}
	}
	return body
}
