package main

import (
	"fmt" 
	"time"
	"math/rand"
	"os"
	"net/http"
)

func main() {

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "The current time is: %s\n", time.Now())
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		fmt.Fprintf(w, "My favorite number is %d", r1.Intn(100))
	})

	http.ListenAndServe(":80", nil)
	
	fmt.Printf("hello, ", os.Args[2])
	fmt.Println("The current time is: ", time.Now())
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	fmt.Println("My favorite number is", r1.Intn(100))
}