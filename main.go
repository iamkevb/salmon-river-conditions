package main

import (
	"fmt"
	"mime"
	"net/http"
	"os"
	"text/template"

	"com.iamkevb.fishing/data"
)

var isDev = false

type PrecipitationViewData struct {
	Dates string
	Rain  string
	Snow  string
}

func main() {
	isDev = len(os.Getenv("DEV")) > 0
	mime.AddExtensionType(".css", "text/css")
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))

	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/temperature.js", handleTemperature)
	http.HandleFunc("/precipitation.js", handlePrecipitation)
	http.HandleFunc("/pressure.js", handlePressure)

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
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := data.Data()

	w.Header().Set("Content-Type", "application/javascript")
	err = tmpl.Execute(w, data)

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

	data := data.Data()
	if isDev {
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	}
	w.Header().Set("Content-Type", "application/javascript")
	err = tmpl.Execute(w, data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handlePressure(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/pressure.tmpl.js")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := data.Data()
	if isDev {
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	}
	w.Header().Set("Content-Type", "application/javascript")
	err = tmpl.Execute(w, data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
