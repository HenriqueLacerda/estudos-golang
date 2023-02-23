package main

import (
	"html/template"
	"net/http"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

type Produto struct {
	Nome       string
	Descricao  string
	Quantidade int
	Preco      float64
}

func main() {

	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{"Produto 1", "descricao 1", 100, 34.50},
		{"Produto 2", "descricao 2", 20, 50.00},
		{"Produto 3", "descricao 3", 400, 87.30},
	}
	temp.ExecuteTemplate(w, "Index", produtos)
}
