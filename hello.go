package main

import (
	"fmt" 
	"time"
	"math/rand"
)

func main() {
	fmt.Printf("hello, world\n")
	fmt.Println("The current time is: ", time.Now())
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	fmt.Println("My favorite number is", r1.Intn(100))
}