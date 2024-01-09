package data

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type WeatherData struct {
	Dates    []time.Time
	Times    []time.Time
	Rain     []float32
	Snow     []float32
	Temps    [][]float32
	MaxHigh  float32
	MinLow   float32
	Pressure []float32
}

type apiData struct {
	Daily  apiDailyData
	Hourly apiHourlyData
}

type apiDailyData struct {
	Time               []string
	Temperature_2m_max []float32
	Temperature_2m_min []float32
	Rain_sum           []float32
	Snowfall_sum       []float32
}

type apiHourlyData struct {
	Time         []string
	Pressure_msl []float32
}

func fetchData(lat float64, lon float64) []byte {
	var apiURL = fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%f&longitude=%f&hourly=pressure_msl&daily=temperature_2m_max,temperature_2m_min,rain_sum,snowfall_sum&timezone=America%%2FNew_York&past_days=7&forecast_days=5", lat, lon)
	response, err := http.Get(apiURL)
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

func fetchWeatherData(lat, lon float64) *WeatherData {
	body := fetchData(lat, lon)
	if len(body) <= 0 {
		return nil
	}
	var data apiData

	err := json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error parsing weather response:", err, string(body))
	}

	temps := [][]float32{}
	dates := []time.Time{}
	var mx float32 = -999
	var mn float32 = 999
	daily := data.Daily
	for i, d := range daily.Time {
		parsed, _ := time.Parse("2006-01-02", d)
		dates = append(dates, parsed)
		t := []float32{daily.Temperature_2m_min[i], daily.Temperature_2m_max[i]}
		mn = min(mn, t[0])
		mx = max(mx, t[1])
		temps = append(temps, t)
	}

	times := []time.Time{}
	hourly := data.Hourly
	for _, t := range hourly.Time {
		parsedTime, _ := time.Parse("2006-01-02T15:04", t)
		times = append(times, parsedTime)
	}

	return &WeatherData{
		Dates:    dates,
		Times:    times,
		Temps:    temps,
		MaxHigh:  mx,
		MinLow:   mn,
		Rain:     daily.Rain_sum,
		Snow:     daily.Snowfall_sum,
		Pressure: hourly.Pressure_msl,
	}
}
