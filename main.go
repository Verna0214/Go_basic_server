package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandle (w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm error: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful.")

	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name is: %s\n", name)
	fmt.Fprintf(w, "Address is: %s\n", address)
}

func helloHandle (w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "hello" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello!")
}

func main () {
	fs := http.FileServer(http.Dir("./static"))

	http.Handle("/", fs)
	http.HandleFunc("/form", formHandle)
	http.HandleFunc("./hello", helloHandle)

	fmt.Println("Server is running on 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}