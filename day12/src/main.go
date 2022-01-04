/**
 * Created by visual studio code
 * @author Iceberry
 * @date 2022-1-3
 * @version 1.0
 * 条件语句select
 * */
package main

import (
	"fmt"
	"time"
)

func main() {
	example1()
	example2()
}

var resChan = make(chan int)

func example2() {
	select {
	case data := <-resChan:
		fmt.Println(data)
	case <-time.After(time.Second * 3):
		fmt.Println("request time out!")
	}
	printSep()
}

func example1() {
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Print("received ", i1, " from c1\n")
	case c2 <- i2:
		fmt.Print("sent ", i2, " to c2\n")
	case i3, ok := (<-c3):
		if ok {
			fmt.Print("received ", i3, " from c3\n")
		} else {
			fmt.Print("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication\n")
	}
	printSep()
}

func printSep() {
	fmt.Println("-----------------------------------")
}
