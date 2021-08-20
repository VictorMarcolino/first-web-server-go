package render

import (
	"html/template"
	"log"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, gohtml string) {
	parseFiles, err := template.ParseFiles("./templates/" + gohtml)
	if err != nil {
		log.Println("Error Parsing file")
		return
	}
	err = parseFiles.Execute(w, nil)
	if err != nil {
		log.Println("Error Executing file")
		return
	}
}
