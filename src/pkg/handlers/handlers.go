package handlers

import (
	"github.com/VictorMarcolino/first-web-server-go/pkg/render"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	log.Println("Request to home page")
	render.RenderTemplate(w, "home.page.gohtml")
}
func About(w http.ResponseWriter, r *http.Request) {
	log.Println("Request to About page")
	render.RenderTemplate(w, "About.page.gohtml")
}
func Info(w http.ResponseWriter, r *http.Request) {
	log.Println("Request to Info page")
	render.RenderTemplate(w, "Info.page.gohtml")
}
