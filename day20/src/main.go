/**
 * Created by visual studio code
 * @author Iceberry
 * @date 2022-1-5
 * @version 1.0
 * 延迟调用(defer)
 * */
package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

func main() {
	example1()
	deferWithClosure()
	test()
	//deferWithError(0)
	deferWithPointer()
	deferWithBigLoop()
	deferWithClosureTrap(2, 0)
	deferWithReturn()
	deferWithNil()
	do()
	do2()
	do3()
}

func do() error {
	res, err := http.Get("http://www.google.com")
	//defer res.Body.Close()
	if res != nil {
		defer res.Body.Close()
	}
	if err != nil {
		return err
	}
	return nil
}

func do2() error {
	f, err := os.Open("book.txt")
	if err != nil {
		return err
	}
	if f != nil {
		//defer f.Close()
		/*
			defer func() {
				if err != f.Close(); err != nil {
					fmt.Println(err)
				}
			}()
		*/
		defer func() {
			if ferr := f.Close(); ferr != nil {
				err = ferr
			}
		}()
	}
	return nil
}

func do3() error {
	f, err := os.Open("book.txt")
	if err != nil {
		return err
	}
	if f != nil {
		defer func() {
			if err := f.Close(); err != nil {
				fmt.Printf("defer close book.txt err %v\n", err)
			}
		}()
	}

	// ..code...

	f, err = os.Open("another-book.txt")
	if err != nil {
		return err
	}
	if f != nil {
		defer func() {
			if err := f.Close(); err != nil {
				fmt.Printf("defer close another-book.txt err %v\n", err)
			}
		}()
	}

	return nil
}

func deferWithNilDefine() {
	var run func() = nil
	defer run()
	fmt.Println("runs")
}

func deferWithNil() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	deferWithNilDefine()
}

func deferWithReturn() (i int) {
	i = 0
	defer func() {
		fmt.Println(i)
	}()
	return 2
}

func deferWithClosureTrap(a, b int) (i int, err error) {
	defer fmt.Printf("first defer err %v\n", err)
	defer func(err error) { fmt.Printf("second defer err %v\n", err) }(err)
	defer func() { fmt.Printf("third defer err %v\n", err) }()
	if b == 0 {
		err = errors.New("divided by zero")
		return
	}
	i = a / b
	return
}

var lock sync.Mutex

func locks() {
	lock.Lock()
	lock.Unlock()
}

func deferLock() {
	lock.Lock()
	defer lock.Unlock()
}
func deferWithBigLoop() {
	func() {
		t1 := time.Now()
		for i := 0; i < 10000; i++ {
			locks()
		}
		elapsed := time.Since(t1)
		fmt.Println("locks elapsed:", elapsed)
	}()
	func() {
		t1 := time.Now()
		for i := 0; i < 10000; i++ {
			deferLock()
		}
		elapsed := time.Since(t1)
		fmt.Println("deferLock elapsed:", elapsed)
	}()
	printSep()
}

func deferWithPointer() {
	x, y := 10, 20
	defer func(i int) {
		fmt.Println("defer:", i, y)
	}(x)
	x += 10
	y += 100
	fmt.Println("x=", x, "y=", y)
}

func deferWithError(x int) {
	defer fmt.Println("a")
	defer fmt.Println("b")
	defer func() {
		fmt.Println(100 / x)
	}()
	defer fmt.Println("c")
}

func test() {
	ts := []Test{{"a"}, {"b"}, {"c"}}
	for _, t := range ts {
		defer t.Close()
	}
	//printSep()
	for _, t := range ts {
		defer Close(t)
	}

	for _, t := range ts {
		t2 := t
		defer t2.Close()
	}
	printSep()

}

func deferWithClosure() {
	var whatever [5]struct{}
	for i := range whatever {
		defer func() {
			fmt.Println(i)
		}()
	}
	printSep()
}

func printSep() {
	fmt.Println("---------------------------------------")
}

func example1() {
	var whatever [5]struct{}
	for i := range whatever {
		defer fmt.Println(i)
	}
	printSep()
}

type Test struct {
	name string
}

func (t *Test) Close() {
	fmt.Println(t.name, "closed")
}

func Close(t Test) {
	t.Close()
}
