package main

import (
	"PersonalWebsite/templates"
	"fmt"
	"net/http"

	"github.com/a-h/templ"
)

func main() {

	componentHello := templates.Hello("Jacob")
	componentProjects := templates.Projects("Jacob")
	componentContact := templates.Contact("Jacob")
	http.Handle("/", templ.Handler(componentHello))
	http.Handle("/projects", templ.Handler(componentProjects))
	http.Handle("/contact", templ.Handler(componentContact))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
