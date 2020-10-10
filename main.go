package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

func connectWithDB() *sql.DB {
	connection := "user=postgres dbname=little_store_go password=1234 host=test-postgres-compose sslmode=disabled"

	db, err := sql.Open("postgres", connection)

	if err != nil {
		panic(err.Error())
	}
	return db
}

type Product struct {
	Name        string
	Description string
	Price       float64
	Quantity    int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	db := connectWithDB()
	defer db.Close()
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
