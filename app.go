package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

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

type Controller interface {
	run(w http.ResponseWriter, r *http.Request)
}

type Home struct {
}

func (h *Home) run(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM product ORDER BY id ASC")
	if err != nil {
		panic(err.Error())
	}
	product := Product{}
	res := []Product{}
	for selDB.Next() {
		var id int
		var name string
		err = selDB.Scan(&id, &name)
		if err != nil {
			panic(err.Error())
		}
		product.Id = id // cart.Sum(1, 5)
		product.Name = name
		res = append(res, product)
	}
	tmpl.ExecuteTemplate(w, "Index", res)
	defer db.Close()
}

type Show struct {
}

func (s *Show) run(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM product WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	product := Product{}
	for selDB.Next() {
		var id int
		var name string
		err = selDB.Scan(&id, &name)
		if err != nil {
			panic(err.Error())
		}
		product.Id = id
		product.Name = name
	}
	tmpl.ExecuteTemplate(w, "Show", product)
	defer db.Close()
}

type CreateNew struct {
}

func (n *CreateNew) run(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

type Edit struct {
}

func (e *Edit) run(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM product WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	product := Product{}
	for selDB.Next() {
		var id int
		var name string
		err = selDB.Scan(&id, &name)
		if err != nil {
			panic(err.Error())
		}
		product.Id = id
		product.Name = name
	}
	tmpl.ExecuteTemplate(w, "Edit", product)
	defer db.Close()
}

type Insert struct {
}

func (i *Insert) run(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		name := r.FormValue("name")
		insForm, err := db.Prepare("INSERT INTO product(name) VALUES(?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name)
		log.Println("INSERT: Name: " + name)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

type Update struct {
}

func (u *Update) run(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		name := r.FormValue("name")
		id := r.FormValue("uid")
		insForm, err := db.Prepare("UPDATE product SET name=? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, id)
		log.Println("UPDATE: Name: " + name)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

type Delete struct{}

func (d *Delete) run(w http.ResponseWriter, r *http.Request) {
	log.Println("pre DELETE")
	db := dbConn()
	product := r.URL.Query().Get("id")
	delForm, err := db.Prepare("DELETE FROM product WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(product)
	log.Println("DELETE")
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func Index(w http.ResponseWriter, r *http.Request) {
	requestPath := r.URL.Path
	home := new(Home)
	show := new(Show)
	cnew := new(CreateNew)
	edit := new(Edit)
	insert := new(Insert)
	update := new(Update)
	delete := new(Delete)
	log.Println(requestPath)
	routes := map[string]Controller{
		"/":       home,
		"/show":   show,
		"/new":    cnew,
		"/edit":   edit,
		"/insert": insert,
		"/update": update,
		"/delete": delete,
	}

	for path, controller := range routes {
		if path == requestPath {
			controller.run(w, r)
			return
		}
	}

	w.WriteHeader(404)
	fmt.Fprint(w, "custom 404")
	return
}

func main() {
	log.Println("Server started on: http://localhost:8080")
	http.HandleFunc("/", Index)
	http.ListenAndServe(":8080", nil)
}
