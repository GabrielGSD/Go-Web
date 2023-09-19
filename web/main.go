package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type Post struct {
	Id    int
	Title string
	Body  string
}

func main() {

	// Definir a rota "/" e a função que será executada
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		post := Post{Id: 1, Title: "Unamed Title", Body: "No Content"}

		if title := r.FormValue("title"); title != "" {
			post.Title = title
		}

		// ParseFiles: ler o arquivo e retornar um template
		t := template.Must(template.ParseFiles("./web/templates/index.html"))
		if err := t.Execute(w, post); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	// Ler a minha porta 8080 e executar o meu servidor
	fmt.Println(http.ListenAndServe(":8080", nil))
}
