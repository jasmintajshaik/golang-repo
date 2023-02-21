package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request succeeded\n\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	age := r.FormValue("age")
	fmt.Fprintf(w, "Your entered details are as below\n")
	fmt.Fprintf(w, "Name: %s\n", name)
	fmt.Fprintf(w, "Age: %s\n", age)
	fmt.Fprintf(w, "Address: %s\n", address)
	fmt.Fprintf(w, "\nThank you! :)")
}
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/about" {
		http.Error(w, "404 Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hi there!, My name is Jasmin. I am learning golang.")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/about", aboutHandler)
	fmt.Printf("Starting server at port 8080\n")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}
