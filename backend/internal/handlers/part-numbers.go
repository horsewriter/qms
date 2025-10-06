package handlers

import (
	"net/http"
	"quality-system/internal/database"
	"quality-system/models"
	"strconv"
	"text/template"

	"github.com/go-chi/chi/v5"
)

func GetPartNumbers(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.GetPartNumbers()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl, err := template.ParseFiles("backend/templates/general_information/part-numbers/part-numbers.gohtml", "backend/templates/general_information/part-numbers/part-number-row.gohtml")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, map[string]interface{}{"PartNumbers": rows})
	}
}

func NewPartNumber(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("backend/templates/general_information/part-numbers/part-number-form.gohtml")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, nil)
	}
}

func CreatePartNumber(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		number := r.FormValue("number")
		newPartNumber := models.PartNumber{Number: number}

		id, err := db.CreatePartNumber(newPartNumber)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl, err := template.ParseFiles("backend/templates/general_information/part-numbers/part-number-row.gohtml")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, models.PartNumber{ID: int(id), Number: number})
	}
}

func EditPartNumber(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))

		partNumber, err := db.GetPartNumberByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl, err := template.ParseFiles("backend/templates/general_information/part-numbers/part-number-form.gohtml")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, map[string]interface{}{"PartNumber": partNumber})
	}
}

func UpdatePartNumber(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))
		number := r.FormValue("number")

		updatedPartNumber := models.PartNumber{ID: id, Number: number}
		err := db.UpdatePartNumber(updatedPartNumber)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl, err := template.ParseFiles("backend/templates/general_information/part-numbers/part-number-row.gohtml")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, updatedPartNumber)
	}
}

func DeletePartNumber(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))

		err := db.DeletePartNumber(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Write([]byte(""))
	}
}

func SearchPartNumbers(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		search := r.FormValue("search")

		rows, err := db.SearchPartNumbers(search)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl, err := template.ParseFiles("backend/templates/general_information/part-numbers/part-number-row.gohtml")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		for _, partNumber := range rows {
			tmpl.Execute(w, partNumber)
		}
	}
}
