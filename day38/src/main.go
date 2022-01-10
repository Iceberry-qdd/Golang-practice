/**
 * Created by visual studio code
 * @author: Iceberry
 * @date: 2022-01-09
 * @version: 1.0
 * 定时器
 */

package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(2 * time.Second)
	t1 := time.Now()
	fmt.Printf("t1:%v\n", t1)
	t2 := <-timer1.C
	fmt.Printf("t2:%v\n", t2)
	printSep()

	// timer2 := time.NewTimer(time.Second)
	// for {
	// 	<-timer2.C
	// 	fmt.Println("时间到")
	// }

	time.Sleep(time.Second)
	fmt.Println("1秒到")
	timer3 := time.NewTimer(2 * time.Second)
	<-timer3.C
	fmt.Println("2秒到")
	timer4 := time.NewTimer(3 * time.Second)
	<-timer4.C
	fmt.Println("3秒到")
	printSep()

	timer5 := time.NewTimer(2 * time.Second)
	go func() {
		<-timer5.C
		fmt.Println("定时器执行了")
	}()
	b := timer5.Stop()
	if b {
		fmt.Println("timer5已经关闭")
	}
	printSep()

	timer6 := time.NewTimer(3 * time.Second)
	timer6.Reset(1 * time.Second)
	fmt.Println(time.Now())
	fmt.Println(<-timer6.C)
	printSep()

	ticker := time.NewTicker(1 * time.Second)
	i := 0
	go func() {
		for {
			i++
			fmt.Println(<-ticker.C)
			if i == 5 {
				ticker.Stop()
			}
		}
	}()
}

func printSep() {
	fmt.Println("--------------------------------------------")
}
