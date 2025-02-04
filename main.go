package main

import (
	"Portfolio/pages"
	"Portfolio/templates"
	"fmt"
	"log"
	"net/http"
)

func main() {

	templates.InitTemplates()
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	http.HandleFunc("/", pages.HomePage)
	http.HandleFunc("/aaa", pages.TableauDeBord)

	fmt.Println("Serveur démarré sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
