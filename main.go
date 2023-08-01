package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/contact", contactHandler)

	fmt.Println("Starting server at http://localhost:8090")
	log.Fatal(http.ListenAndServe(":8090", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "src/templates/pages/home.html")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "src/teplates/pages/contact.html")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "src/teplates/pages/contact.html")
}
