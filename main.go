package main

import (
	"net/http"
	"os"
	_ "expvar"
	"log"
	"fmt"
)


func calculateMonGoQuery(dbName,dbQuery string) (string, error){

	//Possible Queries
	//1. Inserting Documents
	//2. Finding Documents
	//3. Finding Documents using Operators
	    //3.1 $gt / $gte || $lt / $lte - greater than / greater than equals , lesser than / lesser than equals
	    //3.2 $exists - does an attribute exist or not
	    //3.3 $regex - Perl-style pattern matching
	    //3.4 $type - search by type of an element

	//4. Updating Documents
	//5. Removing Documents
	//6. Working with Indexes
	//7. Pipeline Stages
	    //7.1 $project
	    //7.2 $match
	    //7.3 $group
	    //7.4 $sort
	    //7.5 $skip
	    //7.6 $limit
	    //7.7 $unwind
	//8. Aggregation Expressions
	    //8.1 $sum
	    //8.2 $avg
	    //8.3 $min / $max
	    //8.4 $push
	    //8.5 $addToSet
	    //8.6 $first / $last

}

func Convert(w http.ResponseWriter, r *http.Request) {
	var dbName,dbQuery,generateButton string

	if r.Method == "GET" {

		http.ServeFile(w, r, "index.html")

		// Form submitted
		r.ParseMultipartForm(32 << 20)
		//validation :  len(r.Form["username"][0])

		dbName = r.FormValue("dbQuery")
		if len(dbName) == 0 {
			fmt.Fprint(w, "dbName cannot be empty")
			os.Exit(0)
		}
		dbQuery = r.FormValue("dbQuery")
		fmt.Fprint(w , "Entered query is  :", dbQuery)


	} else {
		//POST : generate button action
		generateButton = r.FormValue("generate")

		//check if button is pressed
		if len(generateButton) != 0 {

			ans,err := calculateMonGoQuery(dbName,dbQuery)
			if err != nil {
				log.Println("calculateMonGoQuery Error", err.Error())
			}

			fmt.Fprint(w , ans)
		} else {
			fmt.Fprint(w , "generatebutton is not pressed!")
		}
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
