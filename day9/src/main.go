/**
 * Created by visual studio code
 * @author Iceberry
 * @date 2022-1-2
 * @version 1.0
 * 结构体struct
 * */
package main

import (
	"encoding/json"
	"fmt"
)

type NewInt int

func (m NewInt) sayHello() {
	fmt.Println("Hello,我是一个int.")
}

type MyInt = int

func main() {
	defineNewType()
	printStruct()
	anonymousStruct()
	pointerStruct()
	initStruct()
	structInMemory()
	constructor()
	structMethod()
	structMethodPtrReceiver()
	addMethodForNewInt()
	anonymousField()
	nestedStruct()
	inheritStruct()
	StructAndJson()
	structTag()
}

func structTag() {
	s1 := Stu{
		Id:     1,
		Gender: "女",
		Name:   "Bukky",
	}
	data, err := json.Marshal(s1)
	if err != nil {
		fmt.Println("json marshal failed!")
	}
	fmt.Printf("json str:%s\n", data)
	printSep()
}

func StructAndJson() {
	class := &Class{
		Title:    "101",
		Students: make([]*Stu, 0, 200),
	}
	for i := 0; i < 10; i++ {
		stu := &Stu{
			Name:   fmt.Sprintf("stu%02d", i),
			Gender: "男",
			Id:     i,
		}
		class.Students = append(class.Students, stu)
	}
	data, err := json.Marshal(class)
	if err != nil {
		fmt.Println("json marshal failed!")
		return
	}
	fmt.Printf("json:%s\n", data)
	str := `{"Title":"101","Students":[{"Id":0,"Gender":"男","Name":"stu00"},{"Id":1,"Gender":"男","Name":"stu01"},{"Id":2,"Gender":"男","Name":"stu02"},{"Id":3,"Gender":"男","Name":"stu03"},{"Id":4,"Gender":"男","Name":"stu04"},{"Id":5,"Gender":"男","Name":"stu05"},{"Id":6,"Gender":"男","Name":"stu06"},{"Id":7,"Gender":"男","Name":"stu07"},{"Id":8,"Gender":"男","Name":"stu08"},{"Id":9,"Gender":"男","Name":"stu09"}]}`
	class1 := &Class{}
	err = json.Unmarshal([]byte(str), class1)
	if err != nil {
		fmt.Println("json unmarshal failed!")
		return
	}
	fmt.Printf("class1=%#v\n", class1)
	printSep()
}

func inheritStruct() {
	dog := &Dog{
		feet: 4,
		Animal: &Animal{
			name: "航航",
		},
	}
	dog.wang()
	dog.move()
	printSep()
}

func nestedStruct() {
	user := User{
		name:   "Iceberry",
		gender: "男",
		address: Address{
			province: "广东",
			city:     "深圳",
		},
	}
	fmt.Printf("user=%#v\n", user)
	printSep()
}

func anonymousField() {
	s := student{
		"Iceberry",
		18,
	}
	fmt.Printf("s=%#v\n", s)
	fmt.Println(s.string, s.int)
	printSep()
}

func addMethodForNewInt() {
	var m NewInt
	m.sayHello()
	m = 100
	fmt.Printf("%#v %T\n", m, m)
	printSep()
}

func structMethodPtrReceiver() {
	p := newPerson("Iceberry", "北京", 21)
	fmt.Printf("p=%#v\n", p)
	p.setAge(22)
	fmt.Printf("p=%#v\n", p)
	printSep()
}

func structMethod() {
	p := newPerson("Iceberry", "北京", 21)
	p.Dream()
	printSep()
}

func (p person) Dream() {
	fmt.Printf("住在%s的%d岁的男青年%s的梦想是学好Go语言！\n", p.city, p.age, p.name)
}

func constructor() {
	p := newPerson("Iceberry", "北京", 21)
	fmt.Printf("p=%#v\n", p)
	printSep()
}

func structInMemory() {
	p := person{
		"Iceberry",
		"北京",
		18,
	}
	fmt.Printf("p.name %p\np.city %p\np.age %p\n", &p.name, &p.city, &p.age)
	printSep()
}

func initStruct() {
	p5 := person{
		name: "Icebrry",
		city: "北京",
		age:  18,
	}
	fmt.Printf("p5=%#v\n", p5)
	p6 := &person{
		name: "Icebrry",
		city: "北京",
	}
	fmt.Printf("p6=%#v\n", p6)
	p7 := &person{
		"Icebrry",
		"北京",
		18,
	}
	fmt.Printf("p7=%#v\n", p7)
	printSep()
}

func pointerStruct() {
	var p1 = new(person)
	fmt.Printf("p1: %T\n", p1)
	fmt.Printf("p1=%#v\n", p1)
	var p2 = new(person)
	p2.name = "测试"
	p2.city = "北京"
	p2.age = 18
	fmt.Printf("p2=%#v\n", p2)
	p3 := &person{}
	fmt.Printf("p3: %T\n", p3)
	fmt.Printf("p3=%#v\n", p3)
	p3.name = "博客"
	p3.age = 30
	p3.city = "成都"
	fmt.Printf("p3=%#v\n", p3)
	printSep()
}

func anonymousStruct() {
	var user struct {
		Name string
		Age  int
	}
	user.Name = "Iceberry"
	user.Age = 18
	fmt.Printf("%#v\n", user)
	printSep()
}

func printStruct() {
	var p1 person
	p1.name = "Iceberry"
	p1.city = "北京"
	p1.age = 18
	fmt.Printf("p1=%v\n", p1)
	fmt.Printf("p1=%#v\n", p1)
	printSep()
}

func defineNewType() {
	var a NewInt
	var b MyInt
	fmt.Printf("type of a:%T\n", a)
	fmt.Printf("type of b:%T\n", b)
	printSep()
}

func printSep() {
	fmt.Println("---------------------------------")
}
