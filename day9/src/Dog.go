/**
 * Created by visual studio code
 * @author Iceberry
 * @date 2022-1-2
 * @version 1.0
 * */
package main

import "fmt"

type Dog struct {
	feet int8
	*Animal
}

func (d *Dog) wang() {
	fmt.Printf("%s会汪汪汪~\n", d.name)
}