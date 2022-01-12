/**
 * Created by visual studio code
 * @author: Iceberry
 * @date: 2022-01-11
 * @version: 1.0
 * Hello world服务器
 */
package main

import (
	"fmt"
	"net/http"
)

/*
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World", r.URL.Path)
}*/

/*
type MyHandler struct{}

func main() {
	myHandler := MyHandler{}
	http.Handle("/myHandler", &myHandler)
	http.ListenAndServe(":8080", nil)
}

func (m *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "通过自己创建的处理器请求！")
}
*/

/*
type MyHandler struct{}

func main() {
	myHandler := MyHandler{}
	server := http.Server{
		Addr:        ":8080",
		Handler:     &myHandler,
		ReadTimeout: 2 * time.Second,
	}
	server.ListenAndServe()
}

func (m *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "通过详细配置处理器请求！")
}
*/

/*
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "通过自己创建的多路复用器处理请求", r.URL.Path)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	http.ListenAndServe(":8080", mux)
}
*/

func main() {
	http.HandleFunc("/http", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "测试http协议！")
}
