package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	router "github.com/AchillesKal/go-router"
	_ "github.com/mattn/go-sqlite3"
)

type Product struct {
	Id   int
	Name string
}

var tmpl = template.Must(template.ParseGlob("form/*"))

func dbConn() (db *sql.DB) {
	dbDriver := "sqlite3"
	dbPath := "./data/gocart_db"
	db, err := sql.Open(dbDriver, dbPath)
	statement, _ := db.Prepare("CREATE TABLE IF NOT EXISTS product (id INTEGER PRIMARY KEY, name TEXT)")
	statement.Exec()
	if err != nil {
		panic(err.Error())
	}
	return db
}

func main() {
	appRouter := router.New(index)
	log.Println("Server started on: http://localhost:8080")
	appRouter.Handle("GET", "/", index)
	appRouter.Handle("GET", "/show/:id", showProduct)
	appRouter.Handle("GET", "/new", newProduct)
	appRouter.Handle("GET", "/edit/:id", editProduct)
	appRouter.Handle("POST", "/insert", insertProduct)
	appRouter.Handle("POST", "/update", updateProduct)
	appRouter.Handle("GET", "/delete/:id", deleteProduct)

	http.ListenAndServe(":8080", appRouter)
}
