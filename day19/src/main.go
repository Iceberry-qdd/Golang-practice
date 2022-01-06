/**
 * Created by visual studio code
 * @author Iceberry
 * @date 2022-1-5
 * @version 1.0
 * 闭包和递归
 * */
package main

import (
	"fmt"
)

func main() {
	c := a()
	c()
	c()
	c()
	a()

	f := test()
	f()
	tmp1 := add(10)
	fmt.Println(tmp1(1), tmp1(2))
	tmp2 := add(100)
	fmt.Println(tmp2(1), tmp2(2))

	f1, f2 := test01(10)
	fmt.Println(f1(1), f2(2))
	fmt.Println(f1(3), f2(4))

	var i int = 7
	fmt.Printf("Factorial of %d is %d\n", i, factorial(i))

	var j int
	for j = 0; j < 10; j++ {
		fmt.Printf("%d\n", fibonaci(j))
	}
}

func a() func() int {
	i := 0
	b := func() int {
		i++
		fmt.Println(i)
		return i
	}
	return b
}

func test() func() {
	x := 100
	fmt.Printf("x1 (%p)=%d\n", &x, x)
	return func() {
		fmt.Printf("x2 (%p)=%d\n", &x, x)
	}
}

func add(base int) func(int) int {
	return func(i int) int {
		base += 1
		return base
	}
}

func test01(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func factorial(i int) int {
	if i <= 1 {
		return 1
	}
	return i * factorial(i-1)
}

func fibonaci(i int) int {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return fibonaci(i-1) + fibonaci(i-2)
}
