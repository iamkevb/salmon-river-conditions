package data

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type WeatherData struct {
	Temperatures []Temperature
}

type Temperature struct {
	Date  string    `json:"x"`
	Temps []float32 `json:"y"`
}

type responseData struct {
	Days []responseDay
}
type responseDay struct {
	Datetime string
	Tempmax  float32
	Tempmin  float32
}

var weather WeatherData

func Data() WeatherData {
	fetchWeather()
	return weather
}

func fetchWeather() {
	now := time.Now()
	fourDaysAgo := now.AddDate(0, 0, -4)
	sevenDaysFromNow := now.AddDate(0, 0, 7)
	url := fmt.Sprintf("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/Altmar%%20NY/%s/%s?unitGroup=metric&include=days%%2Chours&key=G8BY72RH6G48WML7P3AGLT46N&contentType=json",
		fourDaysAgo.Format("2006-01-02"),
		sevenDaysFromNow.Format("2006-01-02"))

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var data responseData

	// Unmarshal the JSON data into the struct
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error parsing weather response:", err)
	}

	temps := []Temperature{}
	for _, d := range data.Days {
		t := Temperature{
			Date:  d.Datetime,
			Temps: []float32{d.Tempmin, d.Tempmax},
		}
		temps = append(temps, t)
	}

	weather = WeatherData{
		Temperatures: temps,
	}
}
