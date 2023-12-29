package main

import (
	"fmt"
	"log"
	"mime"
	"net/http"
	"os"
	"text/template"

	"com.iamkevb.fishing/data"
	"github.com/gorilla/mux"
)

var isDev = false

type PrecipitationViewData struct {
	Dates string
	Rain  string
	Snow  string
}

func main() {
	isDev = len(os.Getenv("DEV")) > 0
	r := mux.NewRouter()
	r.Use(loggingMiddleware)
	r.Use(cacheMiddleware)
	r.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))

	mime.AddExtensionType(".css", "text/css")
	// http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))

	r.HandleFunc("/", handleIndex)
	r.HandleFunc("/temperature.js", handleTemperature)
	r.HandleFunc("/precipitation.js", handlePrecipitation)
	r.HandleFunc("/pressure.js", handlePressure)
	r.HandleFunc("/flow.js", handleFlow)

	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Println("Server encountered an error:", err)
	}
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func cacheMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if isDev {
			w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		} else {
			w.Header().Set("Cache-Control", "public, max-age=3600")
		}
		next.ServeHTTP(w, r)
	})
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

func handleFlow(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/flow.tmpl.js")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := data.GetSiteData("04250200")
	if isDev {
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	}
	w.Header().Set("Content-Type", "application/javascript")
	err = tmpl.Execute(w, data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
