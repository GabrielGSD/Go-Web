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
	// Insert Data
	// stmt, err := db.Prepare("INSERT INTO posts(title, body) VALUES(?, ?)")
	// checkErr(err)

	// _, err = stmt.Exec("FullCycle", "FullCycle Rocks!")
	// checkErr(err)

	rows, err := db.Query("SELECT id, title, body FROM posts")
	checkErr(err)

	items := []Post{}

	for rows.Next() {
		item := Post{}
		err = rows.Scan(&item.Id, &item.Title, &item.Body)
		checkErr(err)
		items = append(items, item)
	}

	db.Close()

	// Definir a rota "/" e a função que será executada
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// ParseFiles: ler o arquivo e retornar um template
		t := template.Must(template.ParseFiles("templates/index.html"))
		if err := t.Execute(w, items); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	// Ler a minha porta 8080 e executar o meu servidor
	fmt.Println(http.ListenAndServe(":8080", nil))
}
