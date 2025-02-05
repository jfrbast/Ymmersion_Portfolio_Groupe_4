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
	http.HandleFunc("/portfabio", pages.Portfabio)
	http.HandleFunc("/dashboard", pages.TableauDeBord)
	http.HandleFunc("/team", pages.Team)

	fmt.Println("Serveur démarré sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
