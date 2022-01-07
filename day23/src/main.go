/**
 * Created by visual studio code
 * @author Iceberry
 * @date 2022-1-6
 * @version 1.0
 * 方法
 * */
package main

import "fmt"

type User struct {
	Name  string
	Email string
}

func (u *User) Notify() {
	fmt.Printf("%v:%v\n", u.Name, u.Email)
}

type Data struct {
	x int
}

func (self Data) valueTest() {
	fmt.Printf("Value:%p\n", &self)
}

func (self *Data) pointerTest() {
	fmt.Printf("Pointer:%p\n", self)
}

type PersonD struct {
	id   int
	name string
}

func (p PersonD) valueShowName() {
	fmt.Println(p.name)
}

func (p *PersonD) pointShowName() {
	fmt.Println(p.name)
}

func main() {
	methodExample1()
	methodExample2()
	structTestValue()
	structTestFunc()
}

func structTestFunc() {
	personValue := PersonD{101, "hello world"}
	personValue.valueShowName()
	personValue.pointShowName()
	printSep()

	personPointer := &PersonD{102, "hello golong"}
	personPointer.valueShowName()
	personPointer.pointShowName()
}

func structTestValue() {
	a := 2
	fmt.Println("valueIntTest:", valueIntTest(a))
	b := 5
	fmt.Println("pointerIntTest:", pointerIntTest(&b))
	printSep()
}

func valueIntTest(a int) int {
	return a + 10
}
func pointerIntTest(a *int) int {
	return *a + 10
}

func methodExample2() {
	data := Data{}
	p := &data
	fmt.Printf("Data:%p\n", p)
	data.valueTest()
	data.pointerTest()

	p.valueTest()
	p.pointerTest()
	printSep()
}

func methodExample1() {
	user1 := User{"golong", "golang@golang.com"}
	user1.Notify()
	user2 := User{"go", "go@go.com"}
	user3 := &user2
	user3.Notify()
	printSep()
}

func printSep() {
	fmt.Println("-----------------------------------")
}
