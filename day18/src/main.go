/**
 * Created by visual studio code
 * @author Iceberry
 * @date 2022-1-4
 * @version 1.0
 * 匿名函数
 * */
package main

import (
	"fmt"
	"math"
)

func main() {
	getSqrt := func(a float64) float64 {
		return math.Sqrt(a)
	}
	fmt.Println(getSqrt(4))
	fn := func() {
		fmt.Println("Hello World!")
	}
	fn()
	fns := [](func(x int) int){
		func(x int) int { return x + 1 },
		func(x int) int { return x + 2 },
	}
	fmt.Println(fns[0](100))
	d := struct {
		fn func() string
	}{
		fn: func() string { return "Hello, World!" },
	}
	fmt.Println(d.fn())
	fc := make(chan func() string, 2)
	fc <- func() string { return "Hello, World!" }
	fmt.Println((<-fc)())
}
