package main

import (
	"fmt"
	"log"
	"net/http"
)

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/about" {
		http.Error(w, "404 Requested data not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Requested method not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "About Us\n")
	fmt.Fprintf(w, "We are ghopers\n")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parseform() error: %v", err)
		return
	}
	fmt.Fprintf(w, "Post request successfull\n")

	email := r.FormValue("email")
	quote := r.FormValue("quote")

	fmt.Fprintf(w, "email: %s\n", email)
	fmt.Fprintf(w, "quote: %s\n", quote)
}

func main() {
	// Server config
	server := http.FileServer(http.Dir("./public"))
	// Routes
	http.Handle("/", server)
	http.HandleFunc("/contact", contactHandler)
	http.HandleFunc("/about", aboutHandler)

	fmt.Printf("Server is listening to PORT:8000\n")
	if error := http.ListenAndServe(":8000", nil); error != nil {
		log.Fatal(error)
	}
}
