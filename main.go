package main

import (
	"net/http"
	"os"
	_ "expvar"
	"log"
	"fmt"
	"strings"
)

//calculating collection name, queryType, remaining string
func calculateCollectionName (dbQuery string) (string, string,string){

	var queryType, rem string
	fmt.Println("dbQuery > ",dbQuery)
	a := strings.Split(dbQuery,".")
	fmt.Println("a : > ",a)
	a2:= strings.Split(a[1],".")
	fmt.Println("a : > ",a2)
	a3 := strings.Split(a[2],"(")
	fmt.Println("a : > ",a3)
	a4 := strings.Split(a3[1],"{")
	fmt.Println("a : > ",a4)
	a5 := strings.Split(a4[0],"}")
	fmt.Println("a : > ",a5)
	//getting collection Name
	queryType = strings.Title(a3[0])


	rem = a3[1]
	//if rem contains array //append []int
	if (strings.Contains(rem , "[")) {

	}

	//return collectionName, QueryType
	return a2[0], queryType, rem
}

func calculateMonGoQuery(dbName,dbQuery string) (string, error){

	//Possible Queries
	//1. Inserting Documents

	//MongoEx
	//Input :
	// 1. Db Name
	// 2. Query :  db.ships.insert({name:'USS Enterprise-D',operator:'Starfleet',type:'Explorer',class:'Galaxy',crew:750,codes:[10,11,12]})

	//GoEx
	//Output : session.DB("Test").C("ships").Insert({Name: "USS Enterprise-D", Operator: "Starfleet", Type: "Explorer", Class: "Galaxy", Crew : 750, Codes: []int{10, 11, 12}}

	//getting collectionName, queryType and remaining string to be appended in answer
	collectionName, queryType, rem := calculateCollectionName(dbQuery)

	outputString  := "session.DB(\""+dbName + "\").C(\"" + collectionName + "\")."+queryType + "("+rem

	return outputString,nil

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

	fmt.Println("inside convert ")
	fmt.Println("r >", r.Method)

	var dbName,dbQuery,generateButton string

	if r.Method == "GET" {
		http.ServeFile(w, r, "index.html")

	} else {

		// Form submitted
		r.ParseMultipartForm(32 << 20)
		//validation :  len(r.Form["username"][0])

		dbName = r.FormValue("dbName")
		if len(dbName) == 0 {
			fmt.Fprint(w, "dbName cannot be empty")
			os.Exit(0)
		}
		fmt.Fprintln(w , "Entered dbName is  :", dbName)

		dbQuery = r.FormValue("dbQuery")
		fmt.Fprintln(w , "Entered query is  :", dbQuery)


		ans,err := calculateMonGoQuery(dbName,dbQuery)
		if err != nil {
			log.Println("calculateMonGoQuery Error", err.Error())
		}

		fmt.Fprint(w ,"outputString > "+ ans)



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
			fmt.Println("generatebutton is not pressed!")
			//fmt.lo(w , "generatebutton is not pressed!")
		}
	}
}


func handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")

	var dbName,dbQuery string


	// Form submitted
	r.ParseMultipartForm(32 << 20)
	//validation :  len(r.Form["username"][0])

	dbName = r.FormValue("dbName")
	if len(dbName) == 0 {
		fmt.Fprint(w, "dbName cannot be empty")
		os.Exit(0)
	}
	fmt.Fprint(w , "Entered dbName is  :", dbName)

	dbQuery = r.FormValue("dbQuery")
	fmt.Fprint(w , "Entered query is  :", dbQuery)

}


func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9876"
	}
	http.HandleFunc("/", Convert)
	//http.HandleFunc("/tt", handler)
	log.Println("Listening....on port: " + port)
	http.ListenAndServe(":"+port, nil)
}
