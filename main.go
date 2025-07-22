package main

import (
	"fmt"
	"log"
	"mime"
	"net/http"
	"os"
	"text/template"

	"com.iamkevb.fishing/data"
	"com.iamkevb.fishing/view"
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

	fs := http.FileServer(http.Dir("assets"))
	r.Handle("/favicon.ico", fs)
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fs))

	r.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	r.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))

	r.HandleFunc("/", handleIndex)
	r.HandleFunc("/{code}", handleIndex)

	//Ottawa Fly Fishing Club site
	r.HandleFunc("/api/ottawa", handleOttawa)

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
	code := "04250200" //PINEVILLE
	vars := mux.Vars(r)
	nc, ok := vars["code"]
	if ok {
		code = nc
	}
	fmt.Println("code", code)
	tmpl, err := template.ParseFiles("templates/newlook.tmpl.html", "templates/weather_card.tmpl.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := data.GetSiteData(code)
	if data == nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		viewData := view.NewSiteViewData(data)
		err = tmpl.Execute(w, viewData)
		if err != nil {
			fmt.Printf("Error %s\n", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

// It would be ideal to cache the response for an hour I guess.

func handleOttawa(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "https://ottawaflyfishers.ca")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent) // 204 No Content
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json := data.OttawaWaterData()
	w.Write([]byte(json))
}
