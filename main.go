package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Product struct {
	Id   int
	Name string
}

var tmpl = template.Must(template.ParseGlob("form/*"))

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "root"
	dbName := "gocrud_db"
	uri := dbUser + ":" + dbPass + "@tcp(database:3306)/" + dbName
	fmt.Println(uri)
	db, err := sql.Open(dbDriver, uri)
	if err != nil {
		panic(err.Error())
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
