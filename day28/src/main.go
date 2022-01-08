/**
 * Created by visual studio code
 * @author Iceberry
 * @date 2022-1-7
 * @version 1.0
 * 匿名字段
 * */
package main

import "fmt"

type Person struct {
	name string
	sex  string
	age  int
}

type Student struct {
	Person
	id   int
	addr string
}

type Stu struct {
	Person
	id   int
	addr string
	name string
}

type sex string

type Teacher struct {
	Person
	int
	sex
}

type Policeman struct {
	*Person
	id   int
	addr string
}

func main() {
	student1 := Student{Person{"Iceberry", "man", 20}, 1, "Beijing"}
	fmt.Println(student1)
	student2 := Student{Person: Person{"Iceberry", "man", 20}}
	fmt.Println(student2)
	student3 := Student{Person: Person{name: "Iceberry"}}
	fmt.Println(student3)
	fmt.Println("-----------------------------")
	var stu Stu
	stu.name = "Iceberry"
	fmt.Println(stu)
	stu.Person.name = "BerryIce"
	fmt.Println(stu)
	fmt.Println("-----------------------------")
	teacher1 := Teacher{Person{"Iceberry", "man", 18}, 22, "woman"}
	fmt.Println(teacher1)
	fmt.Println("-----------------------------")
	policeman1 := Policeman{&Person{"Iceberry", "man", 18}, 20, "BerryIce"}
	fmt.Println(policeman1)
	fmt.Println(policeman1.name)
	fmt.Println(policeman1.Person.name)
}
