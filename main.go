package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Listening on port 3000")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	http.ListenAndServe(":3000", nil)

}
