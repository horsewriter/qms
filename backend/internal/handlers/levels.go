package handlers

import (
	"net/http"
	"quality-system/internal/database"
	"quality-system/models"
	"strconv"
	"text/template"

	"github.com/go-chi/chi/v5"
)

func GetLevels(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.GetLevels()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl, err := template.ParseFiles("backend/templates/general_information/levels/levels.gohtml", "backend/templates/general_information/levels/level-row.gohtml")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, map[string]interface{}{"Levels": rows})
	}
}

func NewLevel(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("backend/templates/general_information/levels/level-form.gohtml")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, nil)
	}
}

func CreateLevel(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")
		newLevel := models.Level{Name: name}

		id, err := db.CreateLevel(newLevel)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl, err := template.ParseFiles("backend/templates/general_information/levels/level-row.gohtml")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, models.Level{ID: int(id), Name: name})
	}
}

func EditLevel(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))

		level, err := db.GetLevelByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl, err := template.ParseFiles("backend/templates/general_information/levels/level-form.gohtml")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, map[string]interface{}{"Level": level})
	}
}

func UpdateLevel(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))
		name := r.FormValue("name")

		updatedLevel := models.Level{ID: id, Name: name}
		err := db.UpdateLevel(updatedLevel)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl, err := template.ParseFiles("backend/templates/general_information/levels/level-row.gohtml")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, updatedLevel)
	}
}

func DeleteLevel(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))

		err := db.DeleteLevel(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Write([]byte(""))
	}
}

func SearchLevels(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		search := r.FormValue("search")

		rows, err := db.SearchLevels(search)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl, err := template.ParseFiles("backend/templates/general_information/levels/level-row.gohtml")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		for _, level := range rows {
			tmpl.Execute(w, level)
		}
	}
}
