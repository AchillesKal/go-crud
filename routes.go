package main

import (
	"log"
	"net/http"
	"net/url"
)

func index(w http.ResponseWriter, r *http.Request, params url.Values) {
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
		product.Id = id
		product.Name = name
		res = append(res, product)
	}
	tmpl.ExecuteTemplate(w, "Index", res)
	defer db.Close()
}

func showProduct(w http.ResponseWriter, r *http.Request, params url.Values) {
	db := dbConn()
	nId := params["id"][0]

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

func newProduct(w http.ResponseWriter, r *http.Request, params url.Values) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

func editProduct(w http.ResponseWriter, r *http.Request, params url.Values) {
	db := dbConn()
	nId := params["id"][0]
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

func insertProduct(w http.ResponseWriter, r *http.Request, params url.Values) {
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

func updateProduct(w http.ResponseWriter, r *http.Request, params url.Values) {
	db := dbConn()
	log.Println("UPDATE: method: " + r.Method)

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

func deleteProduct(w http.ResponseWriter, r *http.Request, params url.Values) {
	db := dbConn()
	nId := params["id"][0]
	delForm, err := db.Prepare("DELETE FROM product WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(nId)
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}
