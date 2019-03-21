package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Product struct {
	Id   int
	Name string
}

var tmpl = template.Must(template.ParseGlob("form/*"))

func dbConn() (db *sql.DB) {
	connStr := "host=database user=pqgouser dbname=gocrud_db sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", index)
	r.HandleFunc("/show/{id}", showProduct)
	r.HandleFunc("/new", newProduct)
	r.HandleFunc("/edit/{id}", editProduct)
	r.HandleFunc("/insert", insertProduct)
	r.HandleFunc("/update", updateProduct)
	r.HandleFunc("/delete/{id}", deleteProduct)

	log.Fatal(http.ListenAndServe(":8080", r))
}
