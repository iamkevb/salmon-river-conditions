package data

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type ForecastData struct {
	Date time.Time
	Rain float32
	Snow float32
	High float32
	Low  float32
}

func (f ForecastData) FormattedDate() string {
	loc, _ := time.LoadLocation("America/New_York")
	now := time.Now().In(loc)

	switch {
	case f.Date.YearDay() == now.YearDay():
		return "Today"

	default:
		return f.Date.Format("Monday, January 2")
	}
}

func (f ForecastData) ClassName() string {
	loc, _ := time.LoadLocation("America/New_York")
	now := time.Now().In(loc)
	switch {
	case f.Date.YearDay() < now.YearDay():
		return "past-weather"
	case f.Date.YearDay() == now.YearDay():
		return "current-weather"
	default:
		return "future-weather"
	}
}

type PressureData struct {
	Times    []time.Time
	Pressure []float32
}

type WeatherData struct {
	Forecast []ForecastData
	Pressure PressureData
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
	var apiURL = fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%f&longitude=%f&hourly=pressure_msl&daily=temperature_2m_max,temperature_2m_min,rain_sum,snowfall_sum&timezone=America%%2FNew_York&past_days=3&forecast_days=7", lat, lon)
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

	daily := data.Daily
	forecasts := []ForecastData{}
	for i, d := range daily.Time {
		parsed, _ := time.Parse("2006-01-02", d)
		forecast := ForecastData{
			Date: parsed,
			Rain: daily.Rain_sum[i],
			Snow: daily.Snowfall_sum[i],
			High: daily.Temperature_2m_max[i],
			Low:  daily.Temperature_2m_min[i],
		}
		forecasts = append(forecasts, forecast)
	}

	times := []time.Time{}
	hourly := data.Hourly
	for _, t := range hourly.Time {
		parsedTime, _ := time.Parse("2006-01-02T15:04", t)
		times = append(times, parsedTime)
	}

	pressure := PressureData{
		Times:    times,
		Pressure: hourly.Pressure_msl,
	}

	return &WeatherData{
		Forecast: forecasts,
		Pressure: pressure,
	}
}
