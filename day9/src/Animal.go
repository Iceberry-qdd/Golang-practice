/**
 * Created by visual studio code
 * @author Iceberry
 * @date 2022-1-2
 * @version 1.0
 * */
package main

import "fmt"

 type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s会动！\n", a.name)
}