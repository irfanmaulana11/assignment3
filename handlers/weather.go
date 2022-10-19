package handlers

import (
	"assignment3/models"
	"net/http"
	"text/template"
)

func GetWeatherView(w http.ResponseWriter, r *http.Request) {
	wt := models.Status{
		Water: 1,
		Wind:  2,
	}

	tmpl, err := template.ParseFiles("view.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, wt); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
