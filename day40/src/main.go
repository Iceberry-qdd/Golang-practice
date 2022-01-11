/**
 * Created by visual studio code
 * @author: Iceberry
 * @date: 2022-01-10
 * @version: 1.0
 * 并发安全和锁
 */
package main

import (
	"fmt"
	"sync"
	"time"
)

var x int64
var wg sync.WaitGroup
var lock sync.Mutex

var (
	y      int64
	wg2    sync.WaitGroup
	lock2  sync.Mutex
	rwlock sync.RWMutex
)

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
	fmt.Println("--------------------------------")
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg2.Add(1)
		go write()
	}
	for i := 0; i < 1000; i++ {
		wg2.Add(1)
		go read()
	}

	wg2.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
}

func add() {
	for i := 0; i < 5; i++ {
		lock.Lock()
		x = x + 1
		lock.Unlock()
	}
	wg.Done()
}

func write() {
	rwlock.Lock()
	x = x + 1
	time.Sleep(10 * time.Millisecond)
	rwlock.Unlock()
	wg2.Done()
}

func read() {
	rwlock.RLock()
	time.Sleep(time.Millisecond)
	rwlock.RUnlock()
	wg2.Done()
}
