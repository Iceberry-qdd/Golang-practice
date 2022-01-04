/**
 * Created by visual studio code
 * @author Iceberry
 * @date 2022-1-3
 * @version 1.0
 * 条件语句switch
 * */
package main

import "fmt"

func main() {
	example1()
	example2()
	example3()
	example4()
	example5()
	example6()
}

func example6() {
	var n = 0
	switch {
	case n > 0 && n < 10:
		fmt.Println("i>0 and i<10")
	case n > 10 && n < 20:
		fmt.Println("i>10 and i<20")
	default:
		fmt.Println("def")
	}
	printSep()
}

func example5() {
	var m = 0
	switch m {
	case 0, 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("def")
	}
	printSep()
}

func example4() {
	var j = 0
	switch j {
	case 0:
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("def")
	}
	printSep()
}

func example3() {
	var k = 0
	switch k {
	case 0:
		println("fallthrough")
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("def")
	}
	printSep()
}

func example2() {
	var x interface{}
	switch i := x.(type) {
	case nil:
		fmt.Printf("x的类型：%T\r\n", i)
	case int:
		fmt.Printf("x是int型\n")
	case float64:
		fmt.Printf("x是float64型\n")
	case bool, string:
		fmt.Printf("x是bool或string型\n")
	case func(int) float64:
		fmt.Printf("x是func(int)型\n")
	default:
		fmt.Printf("未知型\n")
	}
	printSep()
}

func example1() {
	var grade string = "B"
	var marks int = 90
	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}

	switch {
	case grade == "A":
		fmt.Printf("优秀！\n")
	case grade == "B", grade == "C":
		fmt.Printf("良好!\n")
	case grade == "D":
		fmt.Printf("良好!\n")
	case grade == "E":
		fmt.Printf("良好!\n")
	default:
		fmt.Printf("良好!\n")
	}
	fmt.Printf("你的等级是%s\n", grade)
	printSep()
}

func printSep() {
	fmt.Println("-----------------------------------")
}
