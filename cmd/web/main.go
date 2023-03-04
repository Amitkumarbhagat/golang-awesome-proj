package main

import (
	"fmt"
	"github.com/amitkumarbhagat/golang-awesome-proj/pkg/config"
	"github.com/amitkumarbhagat/golang-awesome-proj/pkg/handlers"
	"github.com/amitkumarbhagat/golang-awesome-proj/pkg/render"
	"net/http"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {

	}

	app.TemplateCache = tc
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/About", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)

}
