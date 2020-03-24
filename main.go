/*
 * @Author: your name
 * @Date: 2020-03-23 09:42:18
 * @LastEditTime: 2020-03-24 13:05:25
 * @Description:
 */
package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("start...")

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/post/", handleRequest)

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Http server ListenAndServe error: ", err)
	}

	fmt.Println("End...")
}
