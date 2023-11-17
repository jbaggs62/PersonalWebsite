package main

import (
	"PersonalWebsite/templates"
	"fmt"
	"net/http"

	"github.com/a-h/templ"
)

func main() {
	// componets for different pages
	componentHello := templates.Hello("Jacob")
	componentProjects := templates.Projects("Jacob")
	componentContact := templates.Contact("Jacob")

	// routes for different pages
	http.Handle("/", templ.Handler(componentHello))
	http.Handle("/projects", templ.Handler(componentProjects))
	http.Handle("/contact", templ.Handler(componentContact))

	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}
