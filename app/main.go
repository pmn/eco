package main

import (
	_ "eco/app/model"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"os"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	// The index page serves up a basic template. Later this will be a page
	// kept on GitHub pages. Sooo... don't do anything too fancy here. It's just
	// for debugging.
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		panic("Couldn't parse template")
	}
	t.Execute(w, nil)
}

func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		remoteAddr := r.RemoteAddr
		if len(remoteAddr) == 0 {
			remoteAddr = r.Header.Get("x-forwarded-for")
		}
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func main() {
	// Set up the site's routes
	r := mux.NewRouter()
	r.HandleFunc("/", defaultHandler)

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	http.Handle("/", r)
	http.Handle("/static/",
		http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))

	log.Printf("[+] starting eco on port %s", port)
	http.ListenAndServe(":"+port, Log(http.DefaultServeMux))
}
