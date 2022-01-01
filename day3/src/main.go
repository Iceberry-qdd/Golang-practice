/**
 * Created by visual studio code
 * @author Iceberry
 * @date 2022-1-1
 * @version 1.0
 * 变量和常量
 * */
package main

import (
	"fmt"
)

var m = 100

const pi = 3.1415926
const e = 2.7182
const (
	n1 = 100
	n2
	n3
)
const (
	m1 = iota
	m2
	_
	m4
)
const (
	p1 = iota
	p2 = 100
	p3 = iota
	p4
)
const p5 = iota

const(
	_=iota
	KB=1<<(10*iota)
	MB=1<<(10*iota)
	GB=1<<(10*iota)
	TB=1<<(10*iota)
	PB=1<<(10*iota)
)
const(
	a1,a2=iota+1,iota+2
	b1,b2
	c1,c2
)

func main() {
	n := 10
	m := 200
	fmt.Println(m, n)
	x, _ := foo()
	_, y := foo()
	fmt.Printf("x=%d y=%s\n", x, y)
	fmt.Printf("pi=%f e=%f\n", pi, e)
	fmt.Printf("n1=%d n2=%d n3=%d\n", n1, n2, n3)
	fmt.Printf("m1=%d m2=%d m4=%d\n", m1, m2, m4)
	fmt.Printf("p1=%d p2=%d p3=%d p4=%d p5=%d\n", p1, p2, p3, p4, p5)
	fmt.Printf("KB=%d MB=%d GB=%d TB=%d PB=%d\n",KB,MB,GB,TB,PB);
	fmt.Printf("a1=%d a2=%d b1=%d b2=%d c1=%d c2=%d\n",a1,a2,b1,b2,c1,c2)
}

func foo() (int, string) {
	return m, "Iceberry"
}
