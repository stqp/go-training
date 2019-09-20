package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	p, err := strconv.Atoi(price)
	if err != nil {
		fmt.Fprintf(w, "parameter is invalid: %q\n", price)
		return
	}
	if _, ok := db[item]; ok {
		fmt.Fprintf(w, "item is already exists: %q\n", item)
		return
	}
	db[item] = dollars(p)
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	p, err := strconv.Atoi(price)
	if err != nil {
		fmt.Fprintf(w, "parameter is invalid: %q\n", price)
		return
	}
	if _, ok := db[item]; !ok {
		fmt.Fprintf(w, "item donesn't exists: %q\n", item)
		return
	}
	db[item] = dollars(p)
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if _, ok := db[item]; !ok {
		fmt.Fprintf(w, "item donesn't exists: %q\n", item)
		return
	}
	delete(db, item)
}

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}
