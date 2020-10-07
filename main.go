package main

import (
	"database/sql"
	"fmt"
	"go-simple-blog/handler"
	"go-simple-blog/model"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

var tpl *template.Template
var db *sql.DB

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
	handler.SetTemplates(tpl)
	var err error
	db, err = sql.Open("postgres", "postgres://goharu:haruyuki0815@localhost/go_simple_blog?sslmode=disable")
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	model.SetDB(db)
	fmt.Println("Database Connect Success & Parse template")
}

func main() {
	fmt.Println("START")
	http.Handle("/src/", http.StripPrefix("/src/", http.FileServer(http.Dir("src/"))))
	http.HandleFunc("/", articleRoute)
	http.ListenAndServe(":8080", nil)
}

func articleRoute(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	urlSlice := strings.Split(url,"/")
	id,err := strconv.Atoi(urlSlice[len(urlSlice)-1])
	switch {
	case url == "/" && r.Method == http.MethodGet:
		log.Println("Hello This is ArticleIndex")
		handler.ArticleIndex(w, r)
	case err == nil && r.Method == http.MethodGet:
		log.Printf("This is Show Page and Id : %d",id)
		handler.ArticleShow(w,r,id)
	case url == "/articles/new" && r.Method == http.MethodGet:
		log.Print("This is New Page")
		handler.ArticleNew(w,r)
	case url == "/articles/" && r.Method == http.MethodPost:
		log.Println("Hello This is ArticleCreate")
		handler.ArticleCreate(w, r)
	default:
		log.Println("Hello This is PageNotFound")
		handler.PageNotFound(w, r)
	}
	//url := strings.Split(r.URL.Path,"/")
	//id,err := strconv.Atoi(url[len(url)-1])
	//if err != nil {
	//	 panic(err)
	//}
	//fmt.Println(id)
}
