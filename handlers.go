package main

import (
	"log"
	"net/http"
	"text/template"
)

// ViewData struct for response of serach
type ViewData struct {
	Discount int        `json:"discount"`
	Criteria string     `json:"criteria"`
	Products []*Product `json:"products"`
}

// IndexHandler handles the index path
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// Set header
	enableCors(&w)
	w.Header().Set("Content-Type", "text/html")

	// Parse templates and handle error
	tpl, err := template.ParseFiles(
		"templates/layout.gohtml",
		"templates/index.gohtml",
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// Execute templates and handle error
	if err := tpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// SearchHandler handles the search path
func SearchHandler(w http.ResponseWriter, r *http.Request) {
	// Initialize Storage
	db, err := NewStorage(getEnvStorage())
	if err != nil {
		log.Fatal(err)
	}

	// Set header
	enableCors(&w)
	w.Header().Set("Content-Type", "text/html")

	// Validate query search
	q := r.URL.Query().Get("q")
	if q == "" || q == "0" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	// Parse templates and handle error
	funcMap := template.FuncMap{
		"hasDiscount": func(d int) bool {
			return d > 0
		},
		"applyDiscount": func(p int, d int) int {
			return (p * d) / 100
		},
	}
	tpl, err := template.New("layout.gohtml").Funcs(funcMap).ParseFiles(
		"templates/layout.gohtml",
		"templates/search.gohtml",
	)

	// Products
	var products []*Product
	products, err = db.Find(q)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// Excecute templates and handle error
	if err := tpl.Execute(w, ViewData{
		Discount: Campaing(q),
		Criteria: q,
		Products: products,
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
