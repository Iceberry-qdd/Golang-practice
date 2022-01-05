/**
 * Created by visual studio code
 * @author Iceberry
 * @date 2022-1-4
 * @version 1.0
 * 函数定义
 * */
package main

import "fmt"

func main() {
	n, s := test(1, 2, "Iceberry")
	fmt.Println(n, s)
	s1 := test2(func() int { return 100 })
	s2 := format(func(s string, x, y int) string {
		return fmt.Sprintf(s, x, y)
	}, "%d,%d", 10, 20)
	fmt.Println(s1, s2)
}

func test(x, y int, s string) (int, string) {
	n := x + y
	return n, fmt.Sprintf(s, n)
}

func test2(fn func() int) int {
	return fn()
}

type FormatFunc func(s string, x, y int) string

func format(fn FormatFunc, s string, x, y int) string {
	return fn(s, x, y)
}
