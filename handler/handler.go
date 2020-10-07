package handler

import (
	"fmt"
	"go-simple-blog/model"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func SetTemplates(t *template.Template)  {
	tpl = t
	fmt.Println("SetUP OK")
}

func ArticleIndex(w http.ResponseWriter,r *http.Request)  {
	articles,err := model.GetArticleList()
	if err != nil {
		fmt.Println("Error")
	}
	if err := tpl.ExecuteTemplate(w,"index.html",articles); err != nil {
		fmt.Println("Error Occured")
	}
}
func ArticleShow(w http.ResponseWriter,r *http.Request,id int)  {
	article,err := model.GetArticle(id)
	if err != nil {
		w.Write([]byte("Page Not Found"))
	}
	if err := tpl.ExecuteTemplate(w,"show.html",article); err != nil {
		fmt.Println("Error Occured")
		w.Write([]byte("Error Occured"))
	}
}
func ArticleNew(w http.ResponseWriter,r *http.Request)  {
	if err := tpl.ExecuteTemplate(w,"new.html",""); err != nil {
		w.Write([]byte("Error Occured"))
	}
}

func ArticleCreate(w http.ResponseWriter,r *http.Request)  {
	_,err := model.CreateArticle(r.PostFormValue("title"),r.PostFormValue("content"))
	if err != nil {
		log.Println("Error ",err.Error())
	}
	http.Redirect(w,r,"/",303)
}
func PageNotFound(w http.ResponseWriter,r *http.Request)  {
	w.Write([]byte("Page Not Found 404"))
}

