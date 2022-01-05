/**
 * Created by visual studio code
 * @author Iceberry
 * @date 2022-1-4
 * @version 1.0
 * 函数返回值
 * */
package main

import "fmt"

func main() {
	var a, b int = 1, 2
	c := add(a, b)
	sum, avg := calc(a, b)
	fmt.Println(a, b, c, sum, avg)
	x, _ := test()
	fmt.Println(x)
	// s := make([]int, 2)
	// s = test()
	fmt.Println(add(test()))
	//fmt.Println(sum(test()))
	fmt.Println(add2(1, 2))
	fmt.Println(add3(1, 2))
	fmt.Println(add4(1, 2))
}

func add(a, b int) (c int) {
	c = a + b
	return
}

func add2(x, y int) (z int) {
	{
		var z = x + y
		return z
	}
}

func add3(x, y int) (z int) {
	defer func() {
		z += 100
	}()
	z = x + y
	return
}

func add4(x, y int) (z int) {
	defer func() {
		fmt.Printf("defer:%d\n", z)
	}()
	z = x + y
	return z + 200
}

func calc(a, b int) (sum int, avg int) {
	sum = a + b
	avg = (a + b) / 2
	return
}

func test() (int, int) {
	return 1, 2
}

func sum(n ...int) int {
	var x int
	for _, i := range n {
		x += i
	}
	return x
}
