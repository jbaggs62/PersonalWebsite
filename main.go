package main

import (
	"PersonalWebsite/templates"
	"fmt"
	"net/http"

	"github.com/a-h/templ"
)

func main() {

	componentHello := templates.Hello("Jacob")
	componentProjects := templates.Project("Jacob")

	http.Handle("/", templ.Handler(componentHello))
	http.Handle("/projects", templ.Handler(componentProjects))
	http.Handle("/contact", templ.Handler(componentProjects))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
