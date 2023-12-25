package data

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type WeatherData struct {
	Temperatures []Temperature
}

type Temperature struct {
	Date  string    `json:"x"`
	Temps []float32 `json:"y"`
}

type apiData struct {
	Daily apiDailyData
}
type apiDailyData struct {
	Time               []string
	Temperature_2m_max []float32
	Temperature_2m_min []float32
	rain_sum           []float32
	snowfall_sum       []float32
}

var apiURL = "https://api.open-meteo.com/v1/forecast?latitude=43.5097&longitude=-76.0022&hourly=pressure_msl&daily=temperature_2m_max,temperature_2m_min,rain_sum,snowfall_sum&timezone=America%2FNew_York&past_days=5&forecast_days=3"
var weather *WeatherData
var fetching bool = false
var lastFetch time.Time

func Data() WeatherData {
	if weather == nil {
		fetchWeather()
	} else if !fetching && lastFetch.Unix() < time.Now().Add(-1*time.Hour).Unix() {
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
	fetching = true

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

	temps := []Temperature{}
	daily := data.Daily
	for i, date := range daily.Time {
		t := Temperature{
			Date:  date,
			Temps: []float32{daily.Temperature_2m_min[i], daily.Temperature_2m_max[i]},
		}
		temps = append(temps, t)
	}

	weather = &WeatherData{
		Temperatures: temps,
	}
	lastFetch = time.Now()
	fetching = false
}
