package handlers

import (
	"net/http"
	"text/template"

	"quality-system/internal/database"
)

func GeneralInformation(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("backend/templates/index.gohtml", "backend/templates/general_information/general_information.gohtml")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.ExecuteTemplate(w, "content", nil)
	}
}
