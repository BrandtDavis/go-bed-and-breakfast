package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	x, y := 2, 2
	mySum := addValues(x, y)
	fmt.Fprintf(w, "This is the about page, and %d + %d is %d", x, y, mySum)
}

// addValues adds 2 integers, and returns the sum
func addValues(x, y int) int {
	return x + y
}

// main is the main application function
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Starting application on port %s", portNumber)
	http.ListenAndServe(portNumber, nil)
}
