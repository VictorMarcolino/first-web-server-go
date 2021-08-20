package main

import (
	"fmt"
	"github.com/VictorMarcolino/first-web-server-go/pkg/config"
	"github.com/VictorMarcolino/first-web-server-go/pkg/handlers"
	"github.com/VictorMarcolino/first-web-server-go/pkg/render"
	"log"
	"net/http"
)

const host = "0.0.0.0:8080"

// main is the main function
func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)
	fmt.Println(fmt.Sprintf("Staring application on port http://%s", host))
	_ = http.ListenAndServe(host, nil)
}
