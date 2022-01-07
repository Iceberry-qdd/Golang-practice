/**
 * Created by visual studio code
 * @author Iceberry
 * @date 2022-1-6
 * @version 1.0
 * 方法匿名字段
 * */
package main

import "fmt"

type User struct {
	id   int
	name string
}

type Manager struct {
	User
	title string
}

func (self *User) toString() string {
	return fmt.Sprintf("User: %p, %v", self, self)
}

func (self *Manager) toString() string {
	return fmt.Sprintf("Manager: %p, %v", self, self)
}

func main() {
	m := Manager{User{1, "Tom"}, "Admin"}
	fmt.Printf("Manager: %p\n", &m)
	fmt.Println(m.toString())
	fmt.Println(m.User.toString())
}
