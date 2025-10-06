butpackage handlers

import (
	"net/http"
	"quality-system/internal/database"
	"quality-system/models"
	"strconv"
	"text/template"

	"github.com/go-chi/chi/v5"
)

func GetAreas(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.GetAreas()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl, err := template.ParseFiles("backend/templates/general_information/areas/areas.gohtml", "backend/templates/general_information/areas/area-row.gohtml")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, map[string]interface{}{"Areas": rows})
	}
}

func NewArea(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("backend/templates/general_information/areas/area-form.gohtml")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, nil)
	}
}

func CreateArea(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")
		newArea := models.Area{Name: name}

		id, err := db.CreateArea(newArea)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl, err := template.ParseFiles("backend/templates/general_information/areas/area-row.gohtml")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, models.Area{ID: int(id), Name: name})
	}
}

func EditArea(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))

		area, err := db.GetAreaByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl, err := template.ParseFiles("backend/templates/general_information/areas/area-form.gohtml")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, map[string]interface{}{"Area": area})
	}
}

func UpdateArea(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))
		name := r.FormValue("name")

		updatedArea := models.Area{ID: id, Name: name}
		err := db.UpdateArea(updatedArea)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl, err := template.ParseFiles("backend/templates/general_information/areas/area-row.gohtml")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, updatedArea)
	}
}

func DeleteArea(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))

		err := db.DeleteArea(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Write([]byte(""))
	}
}

func SearchAreas(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		search := r.FormValue("search")

		rows, err := db.SearchAreas(search)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl, err := template.ParseFiles("backend/templates/general_information/areas/area-row.gohtml")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		for _, area := range rows {
			tmpl.Execute(w, area)
		}
	}
}
