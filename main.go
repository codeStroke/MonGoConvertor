package main

import (
	"net/http"
	"os"
	_ "expvar"
	"log"
	"fmt"
)

func Convert(w http.ResponseWriter, r *http.Request) {
	var dbQuery string
	if r.Method == "GET" {

		http.ServeFile(w, r, "index.html")

		// Form submitted
		r.ParseMultipartForm(32 << 20)
		//validation :  len(r.Form["username"][0])

		dbName := r.FormValue("dbQuery")
		if len(dbName) == 0 {
			fmt.Fprint(w, "dbName cannot be empty")
			os.Exit(0)
		}
		dbQuery = r.FormValue("dbQuery")
		fmt.Fprint(w , "Entered query is  :", dbQuery)


	} else {
		//POST : generate button action


	}


}
func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "9876"
	}
	http.HandleFunc("/", Convert)
	log.Println("Listening....")
	http.ListenAndServe(":"+port, nil)
}
