/**
 * Created by visual studio code
 * @author Iceberry
 * @date 2022-1-3
 * @version 1.0
 * 循环语句range
 * */
package main

import (
	"fmt"
)

func main() {
	example1()
	example2()
	example3()
}

func example3() {
	s := []int{1, 2, 3, 4, 5}
	for i, v := range s {
		if i == 0 {
			s = s[:3]
			s[2] = 100
		}
		fmt.Println(i, v)
	}
	printSep()
}

func example2() {
	a := [3]int{0, 1, 2}
	for i, v := range a {
		if i == 0 {
			a[1], a[2] = 999, 999
			fmt.Println(a)
		}
		a[i] = v + 100
	}
	fmt.Println(a)
	printSep()
}

func example1() {
	s := "abc"
	for i := range s {
		fmt.Println(s[i])
	}
	for _, c := range s {
		fmt.Println(c)
	}
	for range s {

	}
	m := map[string]int{"a": 1, "b": 2}
	for k, v := range m {
		fmt.Println(k, v)
	}
	printSep()
}

func printSep() {
	fmt.Println("-------------------------------")
}
