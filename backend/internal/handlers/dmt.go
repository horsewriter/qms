package handlers

import (
	"net/http"
	"text/template"

	"quality-system/internal/database"
)

func DMT(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles(
			"backend/templates/index.gohtml",
			"backend/templates/menu.gohtml",
			"backend/templates/dmt/dmt.gohtml",
			"backend/templates/dmt/general-information.gohtml",
			"backend/templates/dmt/material-identification.gohtml",
			"backend/templates/dmt/defect-description.gohtml",
			"backend/templates/dmt/process-analysis.gohtml",
			"backend/templates/dmt/disposition.gohtml",
			"backend/templates/dmt/quality.gohtml",
			"backend/templates/dmt/costs.gohtml",
			"backend/templates/dmt/authorizations.gohtml",
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.ExecuteTemplate(w, "index", nil)
	}
}