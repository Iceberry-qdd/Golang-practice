/**
 * Created by visual studio code
 * @author Iceberry
 * @date 2022-1-6
 * @version 1.0
 * 自定义error
 * */
package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

type PathError struct {
	path       string
	op         string
	createTime string
	message    string
}

func (p *PathError) Error() string {
	return fmt.Sprintf("path=%s \nop=%s \ncreateTime=%s\nmessage=%s", p.path, p.op, p.createTime, p.message)
}

func main() {
	//example1()
	//example2()
	//example3()
	//example4()
	//example5()
	example6()
}

func example6() {
	err := Open("/static/index.html")
	switch v := err.(type) {
	case *PathError:
		fmt.Println("get path error,", v)
	default:
	}
}

func Open(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return &PathError{
			path:       filename,
			op:         "read",
			message:    err.Error(),
			createTime: fmt.Sprintf("%v", time.Now()),
		}
	}
	defer file.Close()
	return nil
}

func example5() {
	area, err := getCircleAreaV2(-5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(area)
	}
}

func getCircleAreaV2(radius float32) (area float32, err error) {
	if radius < 0 {
		err = errors.New("半径不能为负数！")
		return
	}
	area = 3.14 * radius * radius
	return
}

func example4() {
	example3()
	fmt.Println("example4")
}

func example3() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	getCircleArea(-5)
	fmt.Println("这里有没有执行")
}

func example2() {
	getCircleArea(-5)
}

func getCircleArea(radius float32) (area float32) {
	if radius < 0 {
		panic("半径不能为负数！")
	}
	return 3.14 * radius * radius
}

func example1() {
	a := [5]int{0, 1, 2, 3, 4}
	a[1] = 123
	fmt.Println(a)
	index := 10
	a[index] = 10
	fmt.Println(a)
}
