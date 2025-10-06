package main

import (
	"net/http"

	"quality-system/internal/database"
	"quality-system/internal/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	db := database.NewDB()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Homepage
	r.Get("/", handlers.Index(db))

	// Defective Material Tag
	r.Get("/dmt", handlers.DMT(db))

	// General Information (Master Data Management)
	r.Route("/general-information", func(r chi.Router) {
		r.Get("/", handlers.GeneralInformation(db))
		// Employees
		r.Get("/employees", handlers.GetEmployees(db))
		r.Get("/employees/new", handlers.NewEmployee(db))
		r.Post("/employees/new", handlers.CreateEmployee(db))
		r.Get("/employees/edit/{id}", handlers.EditEmployee(db))
		r.Post("/employees/edit/{id}", handlers.UpdateEmployee(db))
		r.Delete("/employees/{id}", handlers.DeleteEmployee(db))
		r.Post("/employees/search", handlers.SearchEmployees(db))
		// Areas
		r.Get("/areas", handlers.GetAreas(db))
		r.Get("/areas/new", handlers.NewArea(db))
		r.Post("/areas/new", handlers.CreateArea(db))
		r.Get("/areas/edit/{id}", handlers.EditArea(db))
		r.Post("/areas/edit/{id}", handlers.UpdateArea(db))
		r.Delete("/areas/{id}", handlers.DeleteArea(db))
		r.Post("/areas/search", handlers.SearchAreas(db))
		// Levels
		r.Get("/levels", handlers.GetLevels(db))
		r.Get("/levels/new", handlers.NewLevel(db))
		r.Post("/levels/new", handlers.CreateLevel(db))
		r.Get("/levels/edit/{id}", handlers.EditLevel(db))
		r.Post("/levels/edit/{id}", handlers.UpdateLevel(db))
		r.Delete("/levels/{id}", handlers.DeleteLevel(db))
		r.Post("/levels/search", handlers.SearchLevels(db))
		// Part Numbers
		r.Get("/part-numbers", handlers.GetPartNumbers(db))
		r.Get("/part-numbers/new", handlers.NewPartNumber(db))
		r.Post("/part-numbers/new", handlers.CreatePartNumber(db))
		r.Get("/part-numbers/edit/{id}", handlers.EditPartNumber(db))
		r.Post("/part-numbers/edit/{id}", handlers.UpdatePartNumber(db))
		r.Delete("/part-numbers/{id}", handlers.DeletePartNumber(db))
		r.Post("/part-numbers/search", handlers.SearchPartNumbers(db))
	})

	http.ListenAndServe(":3000", r)
}
