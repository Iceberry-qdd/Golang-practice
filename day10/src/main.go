/**
 * Created by visual studio code
 * @author Iceberry
 * @date 2022-1-3
 * @version 1.0
 * 条件语句
 * */
package main

import "fmt"

func main() {
	example1()
	example2()
	example3()
	example4()
}

func example4() {
	var a int = 100
	var b int = 200
	if a == 100 {
		if b == 200 {
			fmt.Printf("a的值为100，b的值为200\n")
		}
	}
	fmt.Printf("a的值为：%d\n", a)
	fmt.Printf("b的值为：%d\n", b)
	printSep()
}

func example3() {
	var a int = 100
	if a < 20 {
		fmt.Printf("a小于20\n")
	} else {
		fmt.Printf("a不小于20\n")
	}
	fmt.Printf("a的值为：%d\n", a)
	printSep()
}

func example2() {
	var a int = 10
	if a < 20 {
		fmt.Printf("a小于20\n")
	}
	fmt.Printf("a的值为：%d\n", a)
	printSep()
}

func example1() {
	x := 0
	if n := "abc"; x > 0 {
		fmt.Println(n[2])
	} else if x < 0 {
		fmt.Println(n[1])
	} else {
		fmt.Println(n[0])
	}
	printSep()
}

func printSep() {
	fmt.Println("-----------------------------------")
}
