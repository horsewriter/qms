package handlers

import (
	"net/http"
	"quality-system/internal/database"
	"quality-system/models"
	"strconv"
	"text/template"

	"github.com/go-chi/chi/v5"
)

func GetEmployees(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.GetEmployees()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl, err := template.ParseFiles("backend/templates/general_information/employees/employees.gohtml", "backend/templates/general_information/employees/employee-row.gohtml")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, map[string]interface{}{"Employees": rows})
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
		name := r.FormValue("name")
		role := r.FormValue("role")
		newEmployee := models.Employee{Name: name, Role: role}

		id, err := db.CreateEmployee(newEmployee)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl, err := template.ParseFiles("backend/templates/general_information/employees/employee-row.gohtml")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, models.Employee{ID: int(id), Name: name, Role: role})
	}
}

func EditEmployee(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))

		employee, err := db.GetEmployeeByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
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
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))
		name := r.FormValue("name")
		role := r.FormValue("role")

		updatedEmployee := models.Employee{ID: id, Name: name, Role: role}
		err := db.UpdateEmployee(updatedEmployee)
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
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))

		err := db.DeleteEmployee(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Write([]byte(""))
	}
}

func SearchEmployees(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		search := r.FormValue("search")

		rows, err := db.SearchEmployees(search)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl, err := template.ParseFiles("backend/templates/general_information/employees/employee-row.gohtml")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		for _, employee := range rows {
			tmpl.Execute(w, employee)
		}
	}
}
