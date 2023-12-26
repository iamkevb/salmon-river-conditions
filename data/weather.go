package data

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

type WeatherData struct {
	Dates    []string
	Times    []string
	Rain     []float32
	Snow     []float32
	Temps    [][]float32
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

var apiURL = "https://api.open-meteo.com/v1/forecast?latitude=43.5097&longitude=-76.0022&hourly=pressure_msl&daily=temperature_2m_max,temperature_2m_min,rain_sum,snowfall_sum&timezone=America%2FNew_York&past_days=5&forecast_days=3"
var weather *WeatherData
var lastFetch time.Time
var mutex sync.Mutex

func Data() WeatherData {
	mutex.Lock()
	defer mutex.Unlock()

	if weather == nil {
		fetchWeather()
	} else if lastFetch.Add(time.Hour).Before(time.Now()) {
		go fetchWeather()
	}
	return *weather
}

func fetchWebData() []byte {
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

func fetchSampleData() []byte {
	fmt.Println("LOADING SAMPLE DATA!")
	f, err := os.Open("data/sample.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return []byte{}
	}
	fileContent, err := io.ReadAll(f)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return []byte{}
	}
	return fileContent
}

func fetchData() []byte {
	isDev := len(os.Getenv("DEV")) > 0
	if isDev {
		return fetchSampleData()
	}
	return fetchWebData()
}

func fetchWeather() {
	lastFetch = time.Now()

	body := fetchData()
	if len(body) <= 0 {
		return
	}
	var data apiData

	// Unmarshal the JSON data into the struct
	err := json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error parsing weather response:", err, string(body))
	}

	temps := [][]float32{}
	daily := data.Daily
	for i := range daily.Time {
		t := []float32{daily.Temperature_2m_min[i], daily.Temperature_2m_max[i]}
		temps = append(temps, t)
	}

	times := []string{}
	hourly := data.Hourly
	for _, t := range hourly.Time {
		parsedTime, _ := time.Parse("2006-01-02T15:04", t)
		formatted := parsedTime.Format("January 2, 03:04 PM")
		times = append(times, formatted)
	}

	weather = &WeatherData{
		Dates:    daily.Time,
		Times:    times,
		Temps:    temps,
		Rain:     daily.Rain_sum,
		Snow:     daily.Snowfall_sum,
		Pressure: hourly.Pressure_msl,
	}
}
