package main

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func main() {
	// file paths for the templates
	baseTemplatePath := filepath.Join("src", "templates", "base.html")
	homeTemplatePath := filepath.Join("src", "templates", "pages", "home.html")
	contactTemplatePath := filepath.Join("src", "templates", "pages", "contact.html") // Add a new Template like this:
	aboutTemplatePath := filepath.Join("src", "templates", "pages", "about.html")     // xyzTemplatePath := filepath.Join("src", "templates", "pages", "xyz.html")

	// Load the base template and parse the templates
	templates := template.Must(template.ParseFiles(baseTemplatePath, homeTemplatePath, contactTemplatePath, aboutTemplatePath)) // Add the xyzTemplatePath here

	// Home Page: http://localhost:8090/
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templates.ExecuteTemplate(w, "base.html", nil)
	})

	// Contact Page: http://localhost:8090/contact
	http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		templates.ExecuteTemplate(w, "base.html", nil)
	})

	// About Page: http://localhost:8090/about
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		templates.ExecuteTemplate(w, "base.html", nil)
	})

	// 		New Page: http://localhost:8090/xyz
	//	http.HandleFunc("/xyz", func(w http.ResponseWriter, r *http.Request) {
	//		templates.ExecuteTemplate(w, "base.html", nil)
	//	})

	// Start the server
	http.ListenAndServe(":8090", nil)
}
