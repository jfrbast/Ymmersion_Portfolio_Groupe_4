package pages

import (
	"Portfolio/templates"
	"log"
	"net/http"
)

// Page d'accueil
func HomePage(w http.ResponseWriter, r *http.Request) {
	err := templates.Tpl.ExecuteTemplate(w, "accueil", nil)
	if err != nil {
		log.Println("Erreur template accueil:", err)
		http.Error(w, "Erreur serveur", http.StatusInternalServerError)
	}
}

// Page "Portfabio"
func Portfabio(w http.ResponseWriter, r *http.Request) {
	err := templates.Tpl.ExecuteTemplate(w, "portfabio", nil)
	if err != nil {
		log.Println("Erreur template portfabio:", err)
		http.Error(w, "Erreur serveur", http.StatusInternalServerError)
	}
}

// Page "Tableau de Bord"
func TableauDeBord(w http.ResponseWriter, r *http.Request) {
	err := templates.Tpl.ExecuteTemplate(w, "tableaudebord", nil)
	if err != nil {
		log.Println("Erreur template tableau de bord:", err)
		http.Error(w, "Erreur serveur", http.StatusInternalServerError)
	}
}

// Page "Team"
func Team(w http.ResponseWriter, r *http.Request) {
	err := templates.Tpl.ExecuteTemplate(w, "team", nil)
	if err != nil {
		log.Println("Erreur template team:", err)
		http.Error(w, "Erreur serveur", http.StatusInternalServerError)
	}
}

func Defi2(w http.ResponseWriter, r *http.Request) {

	err := templates.Tpl.ExecuteTemplate(w, "defi2", nil)
	if err != nil {
		log.Println("Erreur lors de l'exécution du template :", err)

		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
	}
}

func Defi1(w http.ResponseWriter, r *http.Request) {
	err := templates.Tpl.ExecuteTemplate(w, "defi_1", nil)
	if err != nil {
		log.Println("Erreur lors de l'écution du template :", err)
	}
}

// Page "All-defis"
func Alldefis(w http.ResponseWriter, r *http.Request) {

	err := templates.Tpl.ExecuteTemplate(w, "alldefis", nil)
	if err != nil {
		log.Println("Erreur lors de l'exécution du template :", err)

		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
	}
}
