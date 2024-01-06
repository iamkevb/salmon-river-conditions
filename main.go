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

	fs := http.FileServer(http.Dir("assets"))
	r.Handle("/favicon.ico", fs)
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fs))

	r.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	r.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))

	r.HandleFunc("/", handleIndex)
	r.HandleFunc("/{code}", handleIndex)

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
	tmpl, err := template.ParseFiles("templates/index.tmpl.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := data.GetSiteData(code)
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
