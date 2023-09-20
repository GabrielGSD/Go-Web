package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
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

	// Definição de rotas
	r := mux.NewRouter()
	// http.FileServer: Onde está os arquivos estáticos
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/{id}/view", ViewHandler)

	// Ler a minha porta 8080 e executar o meu servidor
	// r -> rotas
	fmt.Println(http.ListenAndServe(":8080", r))
}

func GetPostById(id int) Post {
	row := db.QueryRow("SELECT * FROM posts WHERE id = ?", id)
	post := Post{}

	err := row.Scan(&post.Id, &post.Title, &post.Body)
	checkErr(err)

	return post
}

func ListPosts() []Post {
	rows, err := db.Query("SELECT id, title, body FROM posts")
	checkErr(err)

	items := []Post{}

	for rows.Next() {
		item := Post{}
		err = rows.Scan(&item.Id, &item.Title, &item.Body)
		checkErr(err)
		items = append(items, item)
	}

	return items
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// ParseFiles: ler o arquivo e retornar um template
	t := template.Must(template.ParseFiles("templates/layout.html", "templates/list.html"))
	if err := t.Execute(w, ListPosts()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	t := template.Must(template.ParseFiles("templates/layout.html", "templates/view.html"))
	if err := t.Execute(w, GetPostById(id)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
