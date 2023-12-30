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

type PrecipitationViewData struct {
	Dates string
	Rain  string
	Snow  string
}

func main() {

	r := mux.NewRouter()
	r.Use(loggingMiddleware)
	r.Use(cacheMiddleware)

	mime.AddExtensionType(".css", "text/css")
	mime.AddExtensionType(".js", "text/javascript")

	r.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	r.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))

	r.HandleFunc("/", handleIndex)
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
	isDev := len(os.Getenv("DEV")) > 0
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
	data := data.GetSiteData("04250200")
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
	w.Header().Set("Content-Type", "application/javascript")
	err = tmpl.Execute(w, data.WaterData)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
