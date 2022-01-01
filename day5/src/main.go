/**
 * Created by visual studio code
 * @author Iceberry
 * @date 2022-1-1
 * @version 1.0
 * 数组Array
 * */
package main

import (
	"fmt"
	"math/rand"
	"time"
)

//一维全局数组
var arr0 [5]int = [5]int{1, 2, 3}
var arr1 = [5]int{1, 2, 3, 4, 5}
var arr2 = [...]int{1, 2, 3, 4, 5, 6}
var str = [5]string{3: "hello world", 4: "Iceberry"}

//高维全局数组
var arr3 [5][3]int
var arr4 [2][3]int = [...][3]int{{1, 2, 3}, {7, 8, 9}}

func main() {
	fmt.Printf("%d %d %d %d\n", len(arr0), len(arr1), len(arr2), len(str))
	fmt.Println(arr0, arr1, arr2, str)
	oneDimArrayInit()
	fmt.Printf("%d %d\n", len(arr3), len(arr4))
	fmt.Println(arr3, arr4)
	twoDimArrayInit()
	twoDimArrayTravelsal()
	printArr(&arr0)
	sumOfArr()
}

func sumOfArr() {
	rand.Seed(time.Now().Unix())
	var arr [10]int
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(1000)
	}
	var sum int = 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	fmt.Printf("sum of arr is %d.\n", sum)
}

func printArr(arr *[5]int) {
	for i, v := range arr {
		fmt.Printf("a[%d]=%d ", i, v)
	}
	fmt.Println()
}

func twoDimArrayTravelsal() {
	var arr [2][3]int = [...][3]int{{1, 2, 3}, {4, 5, 6}}
	for _, v1 := range arr {
		for _, v2 := range v1 {
			fmt.Printf("%d ", v2)
		}
		fmt.Println()
	}
}

func twoDimArrayInit() {
	a := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	b := [...][2]int{{1, 1}, {2, 2}, {3, 3}}
	fmt.Println(len(a), len(b))
	fmt.Println(a, b)
}

func oneDimArrayInit() {
	a := [3]int{1, 2}
	b := [...]int{1, 2, 3, 4}
	c := [5]int{2: 100, 4: 200} //c[2]=100,c[4]=200,其余为0
	d := [...]struct {
		name string
		age  uint8
	}{
		{"user1", 10},
		{"user2", 20},
	}
	fmt.Println(len(a), len(b), len(c), len(d))
	fmt.Println(a, b, c, d)

}
