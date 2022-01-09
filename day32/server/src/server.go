/**
 * Created by visual studio code
 * @author Iceberry
 * @date 2022-1-8
 * @version 1.0
 * http服务端
 * */
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/go", myHandler)
	http.ListenAndServe("127.0.0.1:8000", nil)
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RemoteAddr, "连接成功")
	fmt.Println("method:", r.Method)
	fmt.Println("url:", r.URL.Path)
	fmt.Println("header:", r.Header)
	fmt.Println("body:", r.Body)
	w.Write([]byte("www.Iceberry.com"))
}
