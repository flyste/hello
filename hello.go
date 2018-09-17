package main

import (

	"time"
	"log"
	"math/rand"

	"net/http"
	"html/template"
)

type PageVariables struct {
	Number1         	int
	Number2				int
	Radius				int
	Time         	string
	Date			string
}

func main() {
	http.HandleFunc("/", HomePage)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func HomePage(w http.ResponseWriter, r *http.Request){

    now := time.Now() // find the time right now
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	r2 := rand.New(s1)
    HomePageVars := PageVariables{ //store the date and time in a struct
      Number1: r1.Intn(900),
	  Number2: r2.Intn(900),
      Time: now.Format("15:04:05"),
	  Date: now.Format("02-01-2006"),
	  Radius: r2.Intn(900)/2,
    }

    t, err := template.ParseFiles("hello.html") //parse the html file homepage.html
    if err != nil { // if there is an error
  	  log.Print("template parsing error: ", err) // log it
  	}
    err = t.Execute(w, HomePageVars) //execute the template and pass it the HomePageVars struct to fill in the gaps
    if err != nil { // if there is an error
  	  log.Print("template executing error: ", err) //log it
  	}
}

