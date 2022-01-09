/**
 * Created by visual studio code
 * @author Iceberry
 * @date 2022-1-8
 * @version 1.0
 * websocket
 * */
package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	go h.run()
	router.HandleFunc("/src", myws)
	if err := http.ListenAndServe("127.0.0.1:8000", router); err != nil {
		fmt.Println("err:", err)
	}
}
