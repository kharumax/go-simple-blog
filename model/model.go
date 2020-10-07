package model

import (
	"database/sql"
	"errors"
	"time"
)

var db *sql.DB

func SetDB(d *sql.DB)  {
	db = d
}

type Article struct {
	Id int
	Title string
	Content string
	Created time.Time
}

func GetArticleList() ([]Article,error)  {
	rows,err := db.Query("SELECT * FROM articles")
	if err != nil {
		return nil,err
	}
	defer rows.Close()
	articles := make([]Article,0)
	for rows.Next() {
		article := Article{}
		err := rows.Scan(&article.Id,&article.Title,&article.Content,&article.Created)
		if err != nil {
			return nil,err
		}
		articles = append(articles,article)
	}
	if err = rows.Err(); err != nil {
		return nil,err
	}
	return articles,nil
}

func GetArticle(id int) (Article,error)  {
	article := Article{}
	row := db.QueryRow("SELECT * FROM articles WHERE id = $1",id)
	err := row.Scan(&article.Id,&article.Title,&article.Content,&article.Created)
	if err != nil {
		return article,err
	}
	return article,nil
}

func CreateArticle(title,content string) (Article,error)  {
	article := Article{}
	article.Title = title
	article.Content = content
	article.Created = time.Now()
	if title == "" || content == "" {
		 return article,errors.New("404 Bad Request")
	}
	err := db.QueryRow("INSERT INTO articles (title,content,created) VALUES ($1,$2,$3) RETURNING id;",title,content,time.Now()).Scan(&article.Id)
	if err != nil {
		return article,errors.New("500. Internal Server Error. "+err.Error())
	}
	return article,nil
}

