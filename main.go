package main

import (
	"Portfolio/pages"
	"Portfolio/templates"
	"fmt"
	"log"
	"net/http"
)

// Fonction principale qui initialise le serveur et définit les routes.
func main() {
	templates.InitTemplates() // Chargement des templates

	// Gestion des fichiers statiques (CSS, images, font)
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))


	// Définition des routes du site

	http.HandleFunc("/challenge6", pages.Challenge6)
	http.HandleFunc("/portfabio", pages.Portfabio)
	http.HandleFunc("/dashboard", pages.TableauDeBord)
	http.HandleFunc("/team", pages.Team)
  http.HandleFunc("/defi2", pages.Defi2)
	http.HandleFunc("/", pages.HomePage)
	http.HandleFunc("/all-defis", pages.Alldefis)

	// Démarrage du serveur sur le port 8080
	fmt.Println("Serveur démarré sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
