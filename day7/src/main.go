/**
 * Created by visual studio code
 * @author Iceberry
 * @date 2022-1-2
 * @version 1.0
 * 指针
 * */
package main

import "fmt"

func main() {
	example()
	getPtrValue()
	ptrArg()
	nilPtr()
	//newExample()
	newPtr()
	makePtr()
	ptrExercise()

}

func ptrExercise() {
	var num int = 0
	fmt.Printf("The address of num is %p\n", &num)
	var ptr *int
	ptr = &num
	*ptr = 1
	fmt.Printf("The value of num is %v\n", num)
	printSep()
}

func makePtr() {
	var b map[string]int
	b = make(map[string]int)
	b["测试"] = 100
	fmt.Println(b)
	printSep()
}

func newPtr() {
	a := new(int)
	b := new(bool)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Println(*a)
	fmt.Println(*b)
	printSep()
}

func newExample() {
	var a *int
	*a = 100
	fmt.Println(*a)
	var b map[string]int
	b["测试"] = 100
	fmt.Println(b)
	printSep()
}

func nilPtr() {
	var p *string
	fmt.Println(p)
	fmt.Printf("p的值是%v\n", p)
	if p != nil {
		fmt.Println("非空")
	} else {
		fmt.Println("空值")
	}
	printSep()

}

func ptrArg() {
	a := 10
	modify1(a)
	fmt.Println(a)
	modify2(&a)
	fmt.Println(a)
	printSep()
}

func modify2(x *int) {
	*x = 100
}

func modify1(x int) {
	x = 100
}

func getPtrValue() {
	a := 10
	b := &a
	fmt.Printf("type of b:%T\n", b)
	c := *b
	fmt.Printf("type of c:%T\n", c)
	fmt.Printf("value of c:%v\n", c)
	printSep()
}

func example() {
	a := 10
	b := &a
	fmt.Printf("a:%d ptr:%p\n", a, &a)
	fmt.Printf("b:%p type:%T\n", b, b)
	printSep()
}

func printSep() {
	fmt.Println("-------------------------------------")
}
