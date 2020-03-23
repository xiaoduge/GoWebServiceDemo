/*
 * @Author: your name
 * @Date: 2020-03-23 09:42:18
 * @LastEditTime: 2020-03-23 16:36:14
 * @Description: In User Settings Edit
 */
package main

import "fmt"

type Post struct {
	Id      int
	Content string
	Author  string
}

func main() {
	fmt.Println("start...")
	post, err := Retrieve(14)
	if err != nil {
		fmt.Println("Retrieve err: ", err)
		panic(err)
	}
	fmt.Println(post)
}
