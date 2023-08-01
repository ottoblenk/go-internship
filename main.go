package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	// Define the data for the templates (e.g., title and year)
	// Define the data for the templates (e.g., title and year)
	data := map[string]string{
		"Title": "My Website",
		"Year":  "2023",
	}

	// Load the base template and parse child templates
	templates := template.Must(template.ParseFiles("base_template.html", "child_template1.html", "child_template2.html"))

	// Handle requests using a handler function
	http.HandleFunc("/page1", func(w http.ResponseWriter, r *http.Request) {
		// Execute child_template1.html using the base template and data
		templates.ExecuteTemplate(w, "base_template.html", data)
	})

	http.HandleFunc("/page2", func(w http.ResponseWriter, r *http.Request) {
		// Execute child_template2.html using the base template and data
		templates.ExecuteTemplate(w, "base_template.html", data)
	})

	// Start the server
	http.ListenAndServe(":8080", nil)

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
