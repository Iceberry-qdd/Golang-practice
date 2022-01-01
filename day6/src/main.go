/**
 * Created by visual studio code
 * @author Iceberry
 * @date 2022-1-1
 * @version 1.0
 * 切片Slice
 * */
package main

import (
	"fmt"
	"strings"
)

//全局切片
var arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
var slice0 []int = arr[2:8]
var slice1 []int = arr[0:6]
var slice2 []int = arr[5:10]
var slice3 []int = arr[0:len(arr)]
var slice4 []int = arr[:len(arr)-1]

func main() {
	testSlice()
	printGlobalSlice()
	initSlice()
	createSliceByMake()
	createSliceDirectly()
	visitSliceByPointer()
	twoDimSlice()
	editSliceItem()
	appendSlice()
	overflowCap()
	reMallocCap()
	copySlice()
	copySlice2()
	traversalSlice()
	resizeSlice()
	stringAndSlice()
	changeCharInString()
	stringWithChinese()
	sliceToString()
	other()

}

func sliceToString() {
	slice := []byte{'I', 'c', 'e', 'b', 'e', 'r', 'r', 'y'}
	str := strings.Replace(strings.Trim(fmt.Sprint(slice), "[]"), " ", ",", -1)
	fmt.Println(str)
	printSep()
}

func other() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	d1 := slice[6:8]
	fmt.Println(d1, len(d1), cap(d1))
	d2 := slice[:6:8]
	fmt.Println(d2, len(d2), cap(d2))
	printSep()
}

func stringWithChinese() {
	str := "你好，世界！hello world！"
	s := []rune(str)
	s[3] = '够'
	s[4] = '浪'
	s[12] = 'g'
	s = s[:14]
	str = string(s)
	fmt.Println(str)
	printSep()
}

func changeCharInString() {
	str := "Hello world"
	s := []byte(str)
	s[6] = 'G'
	s = s[:8]
	s = append(s, '!')
	str = string(s)
	fmt.Println(str)
	printSep()
}

func stringAndSlice() {
	str := "hello world"
	s1 := str[0:5]
	fmt.Println(s1)

	s2 := str[6:]
	fmt.Println(s2)
	printSep()
}

func resizeSlice() {
	var s = []int{1, 3, 4, 5}
	fmt.Printf("slice s: %v,len(s): %v\n", s, len(s))
	b := s[1:2]
	fmt.Printf("slice b : %v , len(b) : %v\n", b, len(b))
	c := b[0:3]
	fmt.Printf("slice c : %v , len(c) : %v\n", c, len(c))
	printSep()
}

func traversalSlice() {
	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice := data[:]
	for index, value := range slice {
		fmt.Printf("slice[%v]=%v\n", index, value)
	}
	printSep()
}

func copySlice2() {
	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("array data : ", data)
	s1 := data[8:]
	s2 := data[:5]
	fmt.Printf("slice s1 : %v\n", s1)
	fmt.Printf("slice s2 : %v\n", s2)
	copy(s2, s1)
	fmt.Printf("copied slice s1 : %v\n", s1)
	fmt.Printf("copied slice s2 : %v\n", s2)
	fmt.Println("last array data : ", data)
	printSep()
}

func copySlice() {
	s1 := []int{1, 2, 3, 4, 5}
	fmt.Printf("slice s1: %v\n", s1)
	s2 := make([]int, 10)
	fmt.Printf("slice s2: %v\n", s2)
	copy(s2, s1)
	fmt.Printf("copied slice s1: %v\n", s1)
	fmt.Printf("copied slice s2: %v\n", s2)
	s3 := []int{1, 2, 3}
	fmt.Printf("slice s3: %v\n", s3)
	s3 = append(s3, s2...)
	fmt.Printf("appended slice s3: %v\n", s3)
	s3 = append(s3, 4, 5, 6)
	fmt.Printf("last slice s3: %v\n", s3)
	printSep()
}

func reMallocCap() {
	s := make([]int, 0, 1)
	c := cap(s)

	for i := 0; i < 50; i++ {
		s = append(s, i)
		if n := cap(s); n > c {
			fmt.Printf("cap:%d->%d\n", c, n)
			c = n
		}
	}
	printSep()
}

func overflowCap() {
	data := [...]int{0, 1, 2, 3, 4, 10: 10}
	s := data[:2:3]
	s = append(s, 100, 200)
	fmt.Println(s, data)
	fmt.Println(&s[0], &data[0])
	printSep()
}

func appendSlice() {
	var slice0 = []int{1, 2, 3}
	fmt.Printf("slice0: %v\n", slice0)
	var slice1 = []int{4, 5, 6}
	fmt.Printf("slice1: %v\n", slice1)
	slice2 := append(slice0, slice1...)
	fmt.Printf("slice2: %v\n", slice2)
	slice3 := append(slice3, 7)
	fmt.Printf("slice3: %v\n", slice3)
	slice4 := append(slice3, 8, 9, 10)
	fmt.Printf("slice4: %v\n", slice4)
	printSep()
}

func editSliceItem() {
	d := [5]struct {
		x int
	}{}
	s := d[:]
	d[1].x = 10
	s[2].x = 20
	fmt.Println(d)
	fmt.Printf("%p,%p\n", &d, &s[0])
	printSep()
}

func twoDimSlice() {
	data := [][]int{
		[]int{1, 2, 3},
		[]int{100, 200},
		[]int{11, 22, 33, 44},
	}
	fmt.Println(data)
	printSep()
}

func visitSliceByPointer() {
	s := []int{0, 1, 2, 3}
	p := &s[2]
	*p += 100
	fmt.Println(s)
	printSep()
}

func createSliceDirectly() {
	s1 := []int{0, 1, 2, 3, 8: 100}
	fmt.Println(s1, len(s1), cap(s1))
	s2 := make([]int, 6, 8)
	fmt.Println(s2, len(s2), cap(s2))
	s3 := make([]int, 6)
	fmt.Println(s3, len(s3), cap(s3))
	printSep()
}

func createSliceByMake() {
	var slice0 []int = make([]int, 10)
	var slice1 []int = make([]int, 10, 10)
	fmt.Printf("make local slice0: %v\n", slice0)
	fmt.Printf("make local slice1: %v\n", slice1)
	printSep()
}

func printGlobalSlice() {
	fmt.Printf("Global: arr %v\n", arr)
	fmt.Printf("Global: slice0 %v\n", slice0)
	fmt.Printf("Global: slice1 %v\n", slice1)
	fmt.Printf("Global: slice2 %v\n", slice2)
	fmt.Printf("Global: slice3 %v\n", slice3)
	fmt.Printf("Global: slice4 %v\n", slice4)
	printSep()
}

func initSlice() {
	arr := [...]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	slice0 := arr[2:8]
	slice1 := arr[0:6]
	slice2 := arr[5:10]
	slice3 := arr[0:len(arr)]
	slice4 := arr[:len(arr)-1]
	fmt.Printf("Local: arr %v\n", arr)
	fmt.Printf("Local: slice0 %v\n", slice0)
	fmt.Printf("Local: slice1 %v\n", slice1)
	fmt.Printf("Local: slice2 %v\n", slice2)
	fmt.Printf("Local: slice3 %v\n", slice3)
	fmt.Printf("Local: slice4 %v\n", slice4)
	printSep()
}

func testSlice() {
	var s1 []int
	if s1 == nil {
		fmt.Println("s1为空!")
	} else {
		fmt.Println("s1不为空!")
	}

	s2 := []int{}
	var s3 []int = make([]int, 0)
	fmt.Println(s1, s2, s3)
	printSep()
}

func printSep() {
	fmt.Println("--------------------------------------")
}
