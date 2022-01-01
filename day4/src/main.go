/**
 * Created by visual studio code
 * @author Iceberry
 * @date 2022-1-1
 * @version 1.0
 * 基本数据类型
 * */
package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	printInteger()
	printFloat()
	printComplex()
	printBool()
	printString()
	printByte()
	printRune()
}

func printRune() {
	str1 := "Iceberry🤡"
	for _, r := range str1 {
		fmt.Printf("%c", r)
	}
	fmt.Print(" ")
	runeStr1 := []rune(str1)
	runeStr1[8] = '❤'
	fmt.Println(string(runeStr1))
}

func printByte() {
	str1 := "Iceberry🤡"
	for i := 0; i < len(str1); i++ {
		fmt.Printf("%c", str1[i])
	}
	fmt.Print(" ")
	byteStr1 := []byte(str1)
	byteStr1[8] = 's'
	fmt.Println(string(byteStr1))

}

func printString() {
	const str1 string = "Ice"
	str2 := "berry🤡"
	fmt.Printf("%s\n", str1+str2)
	str3 := `line1
		line2
		line3
	`
	fmt.Print(str3)
	fmt.Print(len(str1), " ", strings.HasPrefix(str2, "be"), " ", strings.Index(str2, "🤡"), "\n")
	arr := []string{"QDD", "is", "your", "father.\n"}
	fmt.Print(strings.Join(arr, " "))
}

func printBool() {
	const a bool = true
	fmt.Print("a=", a)
	fmt.Printf(" !a=%t\n", !a)
}

func printComplex() {
	const comp64 complex64 = 5 + 2i
	const comp128 complex128 = complex(2, 5)
	fmt.Print("comp64=", real(comp64), "+", imag(comp64), "i\n", "comp128=", comp128, "\n")
}

func printFloat() {
	const maxOfFloat32 float32 = math.MaxFloat32
	const maxOfFloat64 float64 = math.MaxFloat64
	fmt.Print("maxOfFloat32=", maxOfFloat32, "\nmaxOfFloat64=", maxOfFloat64, "\n")
}

func printInteger() {
	const maxOfInt8 int8 = math.MaxInt8
	const maxOfInt16 int16 = math.MaxInt16
	const maxOfInt32 int32 = math.MaxInt32
	const maxOfInt64 int64 = math.MaxInt64
	const maxOfInt int = math.MaxInt
	fmt.Print("maxOfInt8=", maxOfInt8, "\nmaxOfInt16=", maxOfInt16, "\nmaxOfInt32=", maxOfInt32, "\nmaxOfInt64=", maxOfInt64, "\nmaxOfInt=", maxOfInt, "\n")
}
