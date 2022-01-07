/**
 * Created by visual studio code
 * @author Iceberry
 * @date 2022-1-6
 * @version 1.0
 * 方法集
 * */
package main

import "fmt"

type T struct {
	int
}

type S struct {
	T
}

func (t T) testT() {
	fmt.Println("类型T方法集包含全部receiver T方法。")
}

func (t *T) testP() {
	fmt.Println("类型*T方法集包含全部receiver *T方法。")
}

func main() {
	t1 := T{1}
	t2 := &t1
	fmt.Printf("t1 is : %v\n", t1)
	t1.testT()
	printSep()
	fmt.Printf("t2 is : %v\n", t2)
	t1.testT()
	t2.testP()
	printSep()
	s1 := S{T{1}}
	s2 := &s1
	fmt.Printf("s1 is : %v\n", s1)
	s1.testT()
	s1.testP()
	printSep()
	fmt.Printf("s2 is : %v\n", s2)
	s2.testT()
	s2.testP()
}

func printSep() {
	fmt.Println("---------------------------------")
}
