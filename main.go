package main

import (
	"fmt"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "src/templates/pages/home.html")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	fmt.Fprint(w, "<h1>About Us</h1><p>This is our website!</p>")
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)

	fmt.Println("Starting server at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// func main() {
// 	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "Hello!")
// 	})
//
// 	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "<h1> we are not home </h1>")
// 	})
//
// 	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "<h1> we are not home </h1>")
// 	})
//
// 	fmt.Printf("Starting server at port 8080\n")
// 	if err := http.ListenAndServe(":8080", nil); err != nil {
// 		log.Fatal(err)
// 	}
// }
