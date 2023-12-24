package main

import (
	"encoding/json"
	"fmt"
	"mime"
	"net/http"
	"text/template"

	"com.iamkevb.fishing/data"
)

func main() {
	mime.AddExtensionType(".css", "text/css")
	mime.AddExtensionType(".js", "application/javascript")
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))

	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/temperature.js", handleTemperature)
	http.HandleFunc("/precipitation.js", handlePrecipitation)

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

	data := data.Data()
	jsonData, err := json.Marshal(data.Temperatures)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = tmpl.Execute(w, string(jsonData))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handlePrecipitation(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/precipitation.tmpl.js")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// data := data.Data()
	// jsonData, err := json.Marshal(data.Temperatures)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	w.Header().Set("Content-Type", "application/json")
	err = tmpl.Execute(w, nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
