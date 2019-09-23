package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
)

func usage() {
	fmt.Println("# Variant must be uppercase letter (A ~ Z) #\n ")
}

func main() {
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type Calculator struct {
	Val  string
	Item Keyboards
}
type Keyboards struct {
	Items []Key
}
type Key struct {
	Val string
}

var (
	input string
)

func index(w http.ResponseWriter, r *http.Request) {
	u, _ := url.Parse(r.URL.String())
	query := u.Query()
	q := query.Get("q")
	switch q {
	case "=":
		expr, err := Parse(input)
		if err != nil {
			input = "0"
			break
		}
		input = fmt.Sprintf("%f", expr.Eval(Env{}))
	case "c":
		input = "0"
	default:
		input += q
	}

	keys := Keyboards{
		Items: []Key{
			{Val: "0"},
			{Val: "1"},
			{Val: "2"},
			{Val: "3"},
			{Val: "4"},
			{Val: "5"},
			{Val: "6"},
			{Val: "7"},
			{Val: "8"},
			{Val: "9"},
			{Val: "+"},
			{Val: "-"},
			{Val: "*"},
			{Val: "/"},
			{Val: "("},
			{Val: ")"},
			{Val: "pow("},
			{Val: "sqrt("},
			{Val: "sin("},
			{Val: "min["},
			{Val: "max["},
			{Val: ","},
			{Val: "="},
			{Val: "c"},
		},
	}
	cal := Calculator{
		Val:  input,
		Item: keys,
	}

	if err := htmlList().Execute(w, cal); err != nil {
		fmt.Fprintf(w, "fail to template\n")
	}

}

func htmlList() *template.Template {
	body := `
	<html>
	<body>
		<form>
			<table>
				<tr><td><input value="{{.Val}}"></td></tr>
				{{range .Item.Items}}
				<tr><td><input type="submit" name="q" value="{{.Val}}"></td></tr>
				{{end}}
			</table>
		</form>
	</body>
	</html>
	`
	return template.Must(template.New("calculator").Parse(body))
}
