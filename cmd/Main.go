package main

// importing necessary packages
import (
	"fmt"
	"net/http"
)

// Defining handler functions
func OriginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the website!\n")
}

// Register Handler functions
func main() {
	http.HandleFunc("/", OriginHandler)
	//setting up file server
	fs := http.FileServer(http.Dir("static/"))
	//handling request with a path preifix
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	//start the server on port 8080
	http.ListenAndServe(":8080", nil)
}
