package main

import (

	"time"
	"encoding/json"
	"log"
	"math/rand"
	"fmt"
	"io/ioutil"
	"net/http"
	"html/template"
)

type PageVariables struct {
	Number1         	int
	Number2				int
	Radius				int
	Time         	string
	Date			string
	Statusdata			string
	Versiondat		string
	FeatCount 		[]string
	FeatName 		[]string
	FeatVersion 	[]string
	
}

func main() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/status/", ServerHealth)
	http.HandleFunc("/features/", Features)
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

func ServerHealth(w http.ResponseWriter, r *http.Request){

	var healthdata = "data"
	var version = "version"
	
	type ServerHealthResponse struct {
	LLS struct {
		Version          string    `json:"version"`
		BuildDate        time.Time `json:"buildDate"`
		BuildVersion     string    `json:"buildVersion"`
		Branch           string    `json:"branch"`
		Patch            string    `json:"patch"`
		FneBuildVersion  string    `json:"fneBuildVersion"`
		ServerInstanceID string    `json:"serverInstanceID"`
		Database         struct {
			ConnectionCheck string `json:"connectionCheck"`
		} `json:"database"`
	} `json:"LLS"`
}

//	response, err := http.Get("http://localhost:7070/api/1.0/instances/~/health")
//		if err != nil {
//			fmt.Printf("The HTTP request failed with error %s\n", err)
//		} else {
//			data, _ := ioutil.ReadAll(response.Body)
//			fmt.Println(string(data))
//			healthdata = string(data)
//		}
		
	response, err := http.Get("http://localhost:7070/api/1.0/instances/~/health")
	if err != nil {
		fmt.Print(err.Error())
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject ServerHealthResponse
	json.Unmarshal(responseData, &responseObject)
	
	healthdata = string(responseObject.LLS.Database.ConnectionCheck)
	version = string(responseObject.LLS.Version)
	
    now := time.Now() // find the time right now

    HealthPageVars := PageVariables{ //store the date and time in a struct
      Time: now.Format("15:04:05"),
	  Date: now.Format("02-01-2006"),
	  Statusdata: healthdata,
	  Versiondat: version,
	  
    }

    t, err := template.ParseFiles("health.html") //parse the html file homepage.html
    if err != nil { // if there is an error
  	  log.Print("template parsing error: ", err) // log it
  	}
    err = t.Execute(w, HealthPageVars) //execute the template and pass it the HomePageVars struct to fill in the gaps
    if err != nil { // if there is an error
  	  log.Print("template executing error: ", err) //log it
  	}
}

func Features(w http.ResponseWriter, r *http.Request){

	type FeatureStruct struct{
		Name string
		Version string
		Count string
		Used int
		}	
	
	type Feature []struct {
	ID                  int       `json:"id"`
	FeatureVersion      string    `json:"featureVersion"`
	SharedUsed          int       `json:"sharedUsed"`
	OverdraftCount      int       `json:"overdraftCount"`
	AssignedReserved    int       `json:"assignedReserved"`
	OverdraftUsedCount  int       `json:"overdraftUsedCount"`
	FeatureCount        string    `json:"featureCount"`
	MeteredReusable     bool      `json:"meteredReusable"`
	Type                string    `json:"type"`
	ReceivedTime        time.Time `json:"receivedTime"`
	FeatureKind         string    `json:"featureKind"`
	Vendor              string    `json:"vendor"`
	UnassignedReserved  int       `json:"unassignedReserved"`
	FeatureID           string    `json:"featureId"`
	Expiry              string    `json:"expiry"`
	MeteredUndoInterval int       `json:"meteredUndoInterval"`
	FeatureName         string    `json:"featureName"`
	Used                int       `json:"used"`
	Metered             bool      `json:"metered"`
	UncappedOverdraft   bool      `json:"uncappedOverdraft"`
	Reserved            int       `json:"reserved"`
	Concurrent          bool      `json:"concurrent"`
	Uncounted           bool      `json:"uncounted"`
	}
		
	response, err := http.Get("http://localhost:7070/api/1.0/instances/~/features")
	if err != nil {
		fmt.Print(err.Error())
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Feature
	json.Unmarshal(responseData, &responseObject)

	featureslist := make([]FeatureStruct, len(responseObject))
	
	for i := 0; i < len(responseObject); i++ {
			featureslist[i].Name = string(responseObject[i].FeatureName)
			featureslist[i].Version = string(responseObject[i].FeatureVersion)
			featureslist[i].Count = string(responseObject[i].FeatureCount)
			featureslist[i].Used = responseObject[i].Used
			}

    t, err := template.ParseFiles("features.html") //parse the html file homepage.html
    if err != nil { // if there is an error
  	  log.Print("template parsing error: ", err) // log it
  	}
    err = t.Execute(w, featureslist) //execute the template and pass it the HomePageVars struct to fill in the gaps
    if err != nil { // if there is an error
  	  log.Print("template executing error: ", err) //log it
  	}
}



