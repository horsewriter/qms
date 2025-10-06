package handlers

import (
	"net/http"
	"text/template"

	"quality-system/internal/database"
)

func Index(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("backend/templates/index.gohtml", "backend/templates/menu.gohtml")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, nil)
	}
}