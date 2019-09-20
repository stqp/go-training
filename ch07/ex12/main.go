package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

type Item struct {
	Name  string
	Price dollars
}
type Result struct {
	Items []Item
}

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func (db database) list(w http.ResponseWriter, req *http.Request) {
	data := Result{}

	for item, price := range db {
		data.Items = append(data.Items, Item{Name: item, Price: price})
	}
	if err := htmlList().Execute(w, data); err != nil {
		fmt.Fprintf(w, "fail to template\n")
	}

}

func htmlList() *template.Template {
	body := `
	<html>
	<body>
		<table>
			<tr>
				<th>Item</th>
				<th>Price</th>
			</tr>
			{{range .Items}}
			<tr>
				<td>{{.Name}}</td>
				<td>{{.Price}}</td>
			</tr>
			{{end}}
		</table>
	</body>
	</html>
	`
	return template.Must(template.New("tracktable").Parse(body))
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
