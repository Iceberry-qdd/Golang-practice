/**
 * Created by visual studio code
 * @author Iceberry
 * @date 2022-1-8
 * @version 1.0
 * websocket
 * */
package main

type Data struct {
	Ip       string   `json:"ip"`
	User     string   `json:"user"`
	From     string   `json:"from"`
	Type     string   `json:"type"`
	Content  string   `json:"content"`
	UserList []string `json:"user_list"`
}
