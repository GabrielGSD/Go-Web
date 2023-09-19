package main

import (
	"fmt"
	"net/http"
	"text/template"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Post struct {
	Id    int
	Title string
	Body  string
}

var db, err = sql.Open("mysql", "root:root@/go_course?charset=utf8")

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	stmt, err := db.Prepare("INSERT INTO posts(title, body) VALUES(?, ?)")
	checkErr(err)

	_, err = stmt.Exec("FullCycle", "FullCycle Rocks!")
	checkErr(err)

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
