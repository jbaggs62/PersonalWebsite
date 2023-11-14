package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
)

func main() {

	componentHello := hello("Jacob")
	componentProjects := projects("Jacob")

	http.Handle("/", templ.Handler(componentHello))
	http.Handle("/projects", templ.Handler(componentProjects))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
