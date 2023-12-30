package data

import (
	"sync"
	"time"
)

type SiteData struct {
	WaterData   *WaterData
	WeatherData *WeatherData
	fetchTime   time.Time
}

var siteDataMap = make(map[string]*SiteData)
var mutex sync.Mutex

func GetSiteData(code string) *SiteData {
	mutex.Lock()
	defer mutex.Unlock()

	siteData, ok := siteDataMap[code]
	if !ok || siteData.fetchTime.Add(time.Hour).Before(time.Now()) {
		siteData = fetchSiteData(code)
		siteData.fetchTime = time.Now()
	}

	siteDataMap[code] = siteData
	return siteData
}

func fetchSiteData(code string) *SiteData {
	waterData := fetchWaterData(code)
	weatherData := fetchWeatherData(waterData.Latitude, waterData.Longitude)
	return &SiteData{
		WaterData:   waterData,
		WeatherData: weatherData,
	}
}
