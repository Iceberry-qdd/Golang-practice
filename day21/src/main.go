/**
 * Created by visual studio code
 * @author Iceberry
 * @date 2022-1-5
 * @version 1.0
 * 异常处理
 * */
package main

import (
	"errors"
	"fmt"
)

func main() {
	test()
	panicWithChan()
	panicWithDefer()
	recoverExample()
	protectCodesExample(2, 1)
	errorType()
	TryAndCatch()
}

func TryAndCatch() {
	Try(func() {
		panic("test panic")
	}, func(err interface{}) {
		fmt.Println(err)
	})
}

func Try(fun func(), handler func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			handler(err)
		}
	}()
	fun()
}

func errorType() {
	defer func() {
		fmt.Println(recover())
	}()
	switch z, err := div(10, 0); err {
	case nil:
		fmt.Println(z)
	case ErrDivByZero:
		panic(err)
	}
}

var ErrDivByZero = errors.New("division by zero")

func div(x, y int) (int, error) {
	if y == 0 {
		return 0, ErrDivByZero
	}
	return x / y, nil
}

func protectCodesExample(x, y int) {
	var z int
	func() {
		defer func() {
			if recover() != nil {
				z = 0
			}
		}()
		panic("test panic")
		z = x / y
		return
	}()
	fmt.Printf("x/y=%d\n", z)
}

func recoverExample() {
	defer func() {
		fmt.Println(recover())
	}()
	defer recover()
	defer fmt.Println(recover())
	defer func() {
		func() {
			fmt.Println("defer inner")
			recover()
		}()
	}()
	panic("test panic")
}

func panicWithDefer() {
	defer func() {
		fmt.Println(recover())
	}()
	defer func() {
		panic("defer panic")
	}()
	panic("test panic")
}

func panicWithChan() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	var ch chan int = make(chan int, 10)
	close(ch)
	ch <- 1
}

func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err.(string))
		}
	}()
	panic("panic error!")
}
