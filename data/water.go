package data

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

type WaterData struct {
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

func fetchWaterData(usgsCode string) *WaterData {
	bytes := loadWaterData(usgsCode)
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

		waterData := &WaterData{
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
		waterData.Times = times
		waterData.Readings = readings
		return waterData
	}

	return nil
}

func loadWaterData(usgsCode string) []byte {
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
