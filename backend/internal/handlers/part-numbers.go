package handlers

import (
	"context"
	"net/http"
	"text/template"

	"quality-system/internal/database"
	"quality-system/internal/models"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetPartNumbers(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		partNumbers, err := db.GetPartNumbers(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl, err := template.ParseFiles("backend/templates/general_information/part-numbers/part-numbers.gohtml", "backend/templates/general_information/part-numbers/part-number-row.gohtml")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, map[string]interface{}{"PartNumbers": partNumbers})
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
		ctx := context.Background()
		number := r.FormValue("number")
		customer := r.FormValue("customer")
		customerID := r.FormValue("customerID")
		newPartNumber := models.PartNumber{Number: number, Customer: customer, CustomerID: customerID}

		result, err := db.CreatePartNumber(ctx, newPartNumber)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		newPartNumber.ID = result.InsertedID.(primitive.ObjectID)

		tmpl, err := template.ParseFiles("backend/templates/general_information/part-numbers/part-number-row.gohtml")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, newPartNumber)
	}
}

func EditPartNumber(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		id, err := primitive.ObjectIDFromHex(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		partNumbers, err := db.GetPartNumbers(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var partNumber *models.PartNumber
		for _, p := range partNumbers {
			if p.ID == id {
				partNumber = &p
				break
			}
		}

		if partNumber == nil {
			http.Error(w, "Part number not found", http.StatusNotFound)
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
		ctx := context.Background()
		id, err := primitive.ObjectIDFromHex(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		number := r.FormValue("number")
		customer := r.FormValue("customer")
		customerID := r.FormValue("customerID")

		updatedPartNumber := models.PartNumber{ID: id, Number: number, Customer: customer, CustomerID: customerID}
		_, err = db.UpdatePartNumber(ctx, id, updatedPartNumber)
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
		ctx := context.Background()
		id, err := primitive.ObjectIDFromHex(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		_, err = db.DeletePartNumber(ctx, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write([]byte(""))
	}
}

func SearchPartNumbers(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		search := r.FormValue("search")

		partNumbers, err := db.GetPartNumbers(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var filtered []models.PartNumber
		for _, p := range partNumbers {
			if search == "" || contains(p.Number, search) || contains(p.Customer, search) {
				filtered = append(filtered, p)
			}
		}

		tmpl, err := template.ParseFiles("backend/templates/general_information/part-numbers/part-number-row.gohtml")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		for _, partNumber := range filtered {
			tmpl.Execute(w, partNumber)
		}
	}
}
