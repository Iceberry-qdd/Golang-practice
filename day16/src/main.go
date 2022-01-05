/**
 * Created by visual studio code
 * @author Iceberry
 * @date 2022-1-4
 * @version 1.0
 * 函数参数
 * */
package main

import (
	"fmt"
)

func main() {
	var a, b int = 1, 2
	swap(&a, &b)
	fmt.Println(a, b)
	fmt.Println(test("sum:%d", 1, 2, 3))
	s := []int{1, 2, 3}
	res := test2("sum:%d", s...)
	fmt.Println(res)
}

func swap(x, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}

func test(s string, n ...int) string {
	var x int
	for _, i := range n {
		x += i
	}
	return fmt.Sprintf(s, x)
}

func test2(s string, n ...int) string {
	var x int
	for _, i := range n {
		x += i
	}
	return fmt.Sprintf(s, x)
}
