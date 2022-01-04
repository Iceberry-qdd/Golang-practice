/**
 * Created by visual studio code
 * @author Iceberry
 * @date 2022-1-3
 * @version 1.0
 * 循环语句for
 * */
package main

import "fmt"

var s string = "abc"

func main() {
	example1()
	example2()
	example3()
	example4()
}

func example4() {
	var i, j int
	for i = 2; i < 100; i++ {
		for j = 2; j <= (i / j); j++ {
			if i%j == 0 {
				break
			}
		}
		if j > (i / j) {
			fmt.Printf("%d是素数\n", i)
		}
	}
}

func example3() {
	numbers := [6]int{1, 2, 3, 5}
	for i, x := range numbers {
		fmt.Printf("第%d位x的值=%d\n", i, x)
	}
	printSep()
}

func example2() {
	n := len(s)
	for n > 0 && n <= 3 {
		fmt.Println(s[n-1])
		n--
	}
	printSep()
}

func example1() {

	for i, n := 0, len(s); i < n; i++ {
		fmt.Println(s[i])
	}
	printSep()
}

func printSep() {
	fmt.Println("------------------------------------")
}
