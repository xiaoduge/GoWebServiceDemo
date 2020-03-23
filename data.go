/*
 * @Author: your name
 * @Date: 2020-03-23 09:42:12
 * @LastEditTime: 2020-03-23 16:35:52
 * @Description: 提供数据库接口
 */

package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var Db *sql.DB

/**
 * @Author: dcj
 * @Date: 2020-03-23 16:35:55
 * @Description:
 * @Param :
 * @Return:
 */
func init() {
	var err error
	Db, err = sql.Open("postgres", "host=111.229.167.91 port=5432 dbname=dcj user=dcj password=dcj sslmode=disable")
	if err != nil {
		fmt.Println("Open database error: ", err)
		panic(err)
	}
}

/**
 * @Author: dcj
 * @Date: 2020-03-23 16:35:59
 * @Description:
 * @Param :
 * @Return:
 */
func Retrieve(id int) (post Post, err error) {
	post = Post{}
	err = Db.QueryRow("select id, content, author from posts where id = $1",
		id).Scan(&post.Id, &post.Content, &post.Author)
	return
}
