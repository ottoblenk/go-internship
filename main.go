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
	http.HandleFunc("/example", exampleHandler)

	fmt.Println("Starting home at http://localhost:8090")
	fmt.Println("Starting about at http://localhost:8090/about")
	fmt.Println("Starting contact at http://localhost:8090/contact")
	fmt.Println("https://localhost:8090/example")
	log.Fatal(http.ListenAndServe(":8090", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "src/templates/pages/home.html")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "src/templates/pages/about.html")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "src/templates/pages/contact.html")
}

func exampleHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "src/templates/pages/example.html")
}
