/**
 * Created by visual studio code
 * @author: Iceberry
 * @date: 2022-01-10
 * @version: 1.0
 * sync
 */
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func hello() {
	defer wg.Done()
	fmt.Println("Hello Goroutine!")
}

func main() {
	wg.Add(1)
	go hello()
	fmt.Println("main goroutine done!")
	wg.Wait()
}
