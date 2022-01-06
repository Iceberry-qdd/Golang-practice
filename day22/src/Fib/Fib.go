/**
 * Created by visual studio code
 * @author Iceberry
 * @date 2022-1-5
 * @version 1.0
 * 单元测试
 * */
package fib

func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}
