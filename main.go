package main

import (
	"fmt"
	"io"
	"mime"
	"net/http"
	"text/template"
)

func main() {

	mime.AddExtensionType(".css", "text/css")
	mime.AddExtensionType(".js", "application/javascript")
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))

	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/temperature.js", handleTemperature)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server encountered an error:", err)
	}
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.tmpl.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleTemperature(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/temperature.tmpl.js")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// w.Header().Set("Content-Type", "application/javascript")
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func fetchWeather() {
	// URL to make the GET request to
	url := "https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/Altmar%20NY/2023-12-19/2023-12-25?unitGroup=metric&include=days%2Chours&key=G8BY72RH6G48WML7P3AGLT46N&contentType=json"

	// Make the GET request
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return
	}
	defer response.Body.Close()

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Print the response body
	fmt.Println(string(body))
}
