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

func GetEmployees(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		employees, err := db.GetEmployees(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl, err := template.ParseFiles("backend/templates/general_information/employees/employees.gohtml", "backend/templates/general_information/employees/employee-row.gohtml")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, map[string]interface{}{"Employees": employees})
	}
}

func NewEmployee(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("backend/templates/general_information/employees/employee-form.gohtml")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, nil)
	}
}

func CreateEmployee(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		name := r.FormValue("name")
		number := r.FormValue("number")
		newEmployee := models.Employee{Name: name, Number: number}

		result, err := db.CreateEmployee(ctx, newEmployee)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		newEmployee.ID = result.InsertedID.(primitive.ObjectID)

		tmpl, err := template.ParseFiles("backend/templates/general_information/employees/employee-row.gohtml")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, newEmployee)
	}
}

func EditEmployee(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		id, err := primitive.ObjectIDFromHex(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		employees, err := db.GetEmployees(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var employee *models.Employee
		for _, e := range employees {
			if e.ID == id {
				employee = &e
				break
			}
		}

		if employee == nil {
			http.Error(w, "Employee not found", http.StatusNotFound)
			return
		}

		tmpl, err := template.ParseFiles("backend/templates/general_information/employees/employee-form.gohtml")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, map[string]interface{}{"Employee": employee})
	}
}

func UpdateEmployee(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		id, err := primitive.ObjectIDFromHex(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		name := r.FormValue("name")
		number := r.FormValue("number")

		updatedEmployee := models.Employee{ID: id, Name: name, Number: number}
		_, err = db.UpdateEmployee(ctx, id, updatedEmployee)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl, err := template.ParseFiles("backend/templates/general_information/employees/employee-row.gohtml")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, updatedEmployee)
	}
}

func DeleteEmployee(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		id, err := primitive.ObjectIDFromHex(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		_, err = db.DeleteEmployee(ctx, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write([]byte(""))
	}
}

func SearchEmployees(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		search := r.FormValue("search")

		employees, err := db.GetEmployees(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var filtered []models.Employee
		for _, e := range employees {
			if search == "" || contains(e.Name, search) || contains(e.Number, search) {
				filtered = append(filtered, e)
			}
		}

		tmpl, err := template.ParseFiles("backend/templates/general_information/employees/employee-row.gohtml")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		for _, employee := range filtered {
			tmpl.Execute(w, employee)
		}
	}
}
