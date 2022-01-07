/**
 * Created by visual studio code
 * @author Iceberry
 * @date 2022-1-6
 * @version 1.0
 * 表达式
 * */
package main

import "fmt"

type User struct {
	id   int
	name string
}

func (self *User) TestP() {
	fmt.Printf("TestP: %p, %v\n", self, self)
}

func (self User) TestT() {
	fmt.Printf("TestT: %p, %v\n", &self, self)
}

func main() {
	u := User{1, "Tom"}
	u.TestP()

	mValve := u.TestP
	mValve()

	mExpression := (*User).TestP
	mExpression(&u)

	u2 := User{1, "Tom"}
	mValve2 := u.TestT
	mValve2()

	u2.id, u2.name = 2, "Jack"
	u2.TestT()

	mPointer := (*User).TestT
	mPointer(&u)
}
