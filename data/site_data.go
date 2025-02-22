package data

import (
	"sync"
	"time"
)

type SiteData struct {
	WaterData   *WaterData
	WeatherData *WeatherData
	ExtraData   ExtraData
	fetchTime   time.Time
}

var siteDataMap = make(map[string]*SiteData)
var mutex sync.Mutex

func GetSiteData(code string) *SiteData {
	mutex.Lock()
	defer mutex.Unlock()

	loc, _ := time.LoadLocation("America/New_York")
	now := time.Now().In(loc)

	siteData, ok := siteDataMap[code]
	if !ok || siteData.fetchTime.Add(time.Hour).Before(now) {
		siteData = fetchSiteData(code)
		if siteData == nil {
			return nil
		}
		siteData.fetchTime = now
	}

	siteDataMap[code] = siteData
	return siteData
}

func fetchSiteData(code string) *SiteData {
	waterData := fetchWaterData(code)
	if waterData == nil {
		return nil
	}
	weatherData := fetchWeatherData(waterData.Latitude, waterData.Longitude)
	extraData := RenderExtraData(code)
	return &SiteData{
		WaterData:   waterData,
		WeatherData: weatherData,
		ExtraData:   extraData,
	}
}
