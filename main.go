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
	// Gestion des fichiers statiques (CSS, images, font)
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	// Définition des routes du site

	http.HandleFunc("/defi_1", pages.Defi1)
	http.HandleFunc("/portfabio", pages.Portfabio)
	http.HandleFunc("/dashboard", pages.TableauDeBord)
	http.HandleFunc("/team", pages.Team)
	http.HandleFunc("/defi2", pages.Defi2)
	http.HandleFunc("/all-defis", pages.Alldefis)
  	http.HandleFunc("/", pages.HomePage)


	// Démarrage du serveur sur le port 8080
	fmt.Println("Serveur démarré sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
