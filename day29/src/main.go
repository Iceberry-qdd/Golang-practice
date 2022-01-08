/**
 * Created by visual studio code
 * @author Iceberry
 * @date 2022-1-7
 * @version 1.0
 * 接口
 * */
package main

import (
	"fmt"
)

type Sayer interface {
	say()
}

type Mover interface {
	move()
}

type WashingMachine interface {
	wash()
	dry()
}

type animal interface {
	Sayer
	Mover
}

type dryer struct{}

func (d dryer) dry() {
	fmt.Println("甩干")
}

type haier struct {
	dryer
}

func (h haier) wash() {
	fmt.Println("haier洗衣")
}

type dog struct{}

func (d dog) say() {
	fmt.Println("汪汪汪")
}

func (d dog) move() {
	fmt.Println("狗会动")
}

type cat struct{}

func (c cat) say() {
	fmt.Println("喵喵喵")
}

type pig struct {
	name string
}

func (p pig) say() {
	fmt.Println("哄哄哄")
}

func (p pig) move() {
	fmt.Println("猪会动")
}

func main() {
	example1()
	example2()
	example3()
	example4()
	example5()
	example6()
	example7()
	example8()
}

func example8() {
	var x interface{}
	x = "Iceberry"
	value, ok := x.(string)
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("类型断言失败")
	}
	justifyType(x)
}

func justifyType(x interface{}) {
	switch value := x.(type) {
	case string:
		fmt.Printf("x is a string,value is %v\n", value)
	case int:
		fmt.Printf("x is a int is %v\n", value)
	case bool:
		fmt.Printf("x is a bool is %v\n", value)
	default:
		fmt.Println("unsupport type！")
	}
}

func example7() {
	var studentInfo = make(map[string]interface{})
	studentInfo["name"] = "Iceberry"
	studentInfo["age"] = 22
	studentInfo["married"] = false
	fmt.Println(studentInfo)
	printSep()
}

func example6() {
	show("Iceberry")
	show([]int{1, 2, 3})
	show(map[string]float32{"apple": 36.25, "banana": 36.87})
	printSep()
}

func show(a interface{}) {
	fmt.Printf("type: %T value: %v\n", a, a)
}

func example5() {
	var x interface{}
	str := "Iceberry"
	x = str
	fmt.Printf("type: %T value: %v\n", x, x)
	i := 100
	x = i
	fmt.Printf("type:%T value:%v\n", x, x)
	b := true
	x = b
	fmt.Printf("type:%T value:%v\n", x, x)
	printSep()
}

func example4() {
	var x animal
	x = pig{name: "花花"}
	x.move()
	x.say()
}

func example3() {
	var w WashingMachine
	var h = haier{}
	w = h
	w.wash()
	w.dry()
	printSep()
}

func example2() {
	var x Mover
	var wangcai = dog{}
	x = wangcai
	x.move()
	var fugui = &dog{}
	x = fugui
	x.move()
	printSep()
}

func example1() {
	var x Sayer
	a := cat{}
	b := dog{}
	x = a
	x.say()
	x = b
	x.say()
	printSep()
}

func printSep() {
	fmt.Println("---------------------------------")
}
