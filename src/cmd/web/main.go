package main

import (
	"github.com/VictorMarcolino/first-web-server-go/pkg/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	http.HandleFunc("/info", handlers.Info)
	log.Println("access http://0.0.0.0:8080")
	_ = http.ListenAndServe("0.0.0.0:8080", nil)
}
