package main

import (
	"net/http"
    "log"
    "html/template"
    // "bytes"

	"github.com/gorilla/mux"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
    // var doc []byte
    t, err := template.ParseGlob("templates/*.html")

    err = t.ExecuteTemplate(w, "index.html", nil)
    if err != nil {
        log.Fatal(err)
    }

	// w.Write(doc)
}



func main() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
    r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public/"))))
	r.HandleFunc("/", YourHandler)

	// Bind to a port and pass our router in
    listen := ":8000"
    log.Printf("Starting server on %s\n", listen)
	http.ListenAndServe(listen, r)
}
