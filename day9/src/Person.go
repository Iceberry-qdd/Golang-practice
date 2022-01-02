/**
 * Created by visual studio code
 * @author Iceberry
 * @date 2022-1-2
 * @version 1.0
 * */
package main

type person struct {
	name string
	city string
	age  int8
}

func newPerson(name, city string, age int8) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}
func (p *person) setAge(age int8) {
	p.age = age
}
