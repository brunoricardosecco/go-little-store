package main

import (
	"html/template"
	"net/http"
)

type Product struct {
	Name        string
	Description string
	Price       float64
	Quantity    int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)

	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	products := []Product{
		{Name: "T-Shirt", Description: "Black T-Shirt", Price: 20, Quantity: 2},
		{Name: "Boot", Description: "Black Boot", Price: 89, Quantity: 3},
		{Name: "HeadPhone", Description: "Headphone test", Price: 250.99, Quantity: 30},
	}

	temp.ExecuteTemplate(w, "index", products)
}
