/**
 * Created by visual studio code
 * @author: Iceberry
 * @date: 2022-01-09
 * @version: 1.0
 * select
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	output1 := make(chan string)
	output2 := make(chan string)
	go test1(output1)
	go test2(output2)

	select {
	case s1 := <-output1:
		fmt.Println("s1=", s1)
	case s2 := <-output2:
		fmt.Println("s2=", s2)
	}
	printSep()

	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)
	go func() {
		int_chan <- 1
	}()
	go func() {
		string_chan <- "hello"
	}()
	select {
	case value := <-int_chan:
		fmt.Println("int:", value)
	case value := <-string_chan:
		fmt.Println("string:", value)
	}
	fmt.Println("main结束")
	printSep()

	output3 := make(chan string, 10)
	go write(output3)
	for s := range output3 {
		fmt.Println("res:", s)
		time.Sleep(time.Second)
	}
}

func write(output3 chan string) {
	for {
		select {
		case output3 <- "hello":
			fmt.Println("write hello")
		default:
			fmt.Println("channel full")
		}
		time.Sleep(time.Millisecond * 500)
	}
}

func printSep() {
	fmt.Println("-------------------------------------------")
}

func test2(output2 chan string) {
	time.Sleep(time.Second * 2)
	output2 <- "test2"
}

func test1(output1 chan string) {
	time.Sleep(time.Second * 5)
	output1 <- "test1"
}
