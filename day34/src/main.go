/**
 * Created by visual studio code
 * @author Iceberry
 * @date 2022-1-8
 * @version 1.0
 * goroutine
 * */
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	// for i := 0; i < 10; i++ {
	// 	wg.Add(1)
	// 	go hello(i)
	// }
	// wg.Wait()

	// fmt.Println("main goroutine done!")
	// time.Sleep(time.Second)

	go func() {
		i := 0
		for {
			i++
			fmt.Printf("new goroutine:i=%d\n", i)
			time.Sleep(time.Second)
		}
	}()
	i := 0
	for {
		i++
		fmt.Printf("main goroutine: i=%d\n", i)
		time.Sleep(time.Second)
		if i == 2 {
			break
		}
	}
}

func hello(i int) {
	defer wg.Done()
	fmt.Println("Hello Goroutine!", i)
}
