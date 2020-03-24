/*
 * @Author: your name
 * @Date: 2020-03-23 09:42:12
 * @LastEditTime: 2020-03-24 11:18:56
 * @Description: 提供数据库接口
 */

package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "host=111.229.167.91 port=5432 dbname=dcj user=dcj password=dcj sslmode=disable")
	if err != nil {
		fmt.Println("Open database error: ", err)
		panic(err)
	}
}

func Retrieve(id int) (post Post, err error) {
	post = Post{}
	err = Db.QueryRow("select id, content, author from posts where id = $1",
		id).Scan(&post.Id, &post.Content, &post.Author)
	return
}

func (post *Post) Create() (err error) {
	statement := "insert into posts (content, author) values ($1, $2) returning id"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	err = stmt.QueryRow(post.Content, post.Author).Scan(post.Id)
	return
}

func (post *Post) Update() (err error) {
	_, err = Db.Exec("update posts set content = $1, author = $2 where id = $3", post.Content, post.Author, post.Id)
	return
}

func (post *Post) Delete() (err error) {
	_, err = Db.Exec("Delete from posts where id = $1", post.Id)
	return
}
