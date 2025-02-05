package pages

import (
	"Portfolio/templates"
	"log"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {

	err := templates.Tpl.ExecuteTemplate(w, "accueil", nil)
	if err != nil {
		log.Println("Erreur lors de l'exécution du template :", err)

		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
	}
}
func Portfabio(w http.ResponseWriter, r *http.Request) {

	err := templates.Tpl.ExecuteTemplate(w, "portfabio", nil)
	if err != nil {
		log.Println("Erreur lors de l'exécution du template :", err)

		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
	}
}

func TableauDeBord(w http.ResponseWriter, r *http.Request) {

	err := templates.Tpl.ExecuteTemplate(w, "tableaudebord", nil)
	if err != nil {
		log.Println("Erreur lors de l'exécution du template :", err)

		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
	}
}