/**
 * Created by visual studio code
 * @author: Iceberry
 * @date: 2022-01-09
 * @version: 1.0
 * channel
 */

package main

import (
	"fmt"
)

func main() {
	var ch chan int
	fmt.Println(ch)
	printSep()

	ch2 := make(chan int)
	go func() {
		ch2 <- 10
	}()
	x := <-ch2
	fmt.Println(x)
	close(ch2)
	printSep()

	ch3 := make(chan int, 1)
	ch3 <- 10
	fmt.Println("发送成功", <-ch3)
	close(ch3)
	printSep()

	ch4 := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch4 <- i
		}
		close(ch4)
	}()

	for {
		if data, ok := <-ch4; ok {
			fmt.Println(data)
		} else {
			break
		}
	}
	printSep()

	ch5 := make(chan int)
	ch6 := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			ch5 <- i
		}
		close(ch5)
	}()
	go func() {
		for {
			i, ok := <-ch5
			if !ok {
				break
			}
			ch6 <- i * i
		}
		close(ch6)
	}()

	for i := range ch6 {
		fmt.Println(i)
	}
	printSep()

	ch7 := make(chan int)
	ch8 := make(chan int)
	go counter(ch7)
	go squarer(ch8, ch7)
	printer(ch8)
}

func printer(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}

func squarer(out chan<- int, in <-chan int) {
	for i := range in {
		out <- i * i
	}
	close(out)
}

func counter(out chan<- int) {
	for i := 0; i < 100; i++ {
		out <- i
	}
	close(out)
}

func printSep() {
	fmt.Println("-------------------------------------")
}
