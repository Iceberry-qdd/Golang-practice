/**
 * Created by visual studio code
 * @author Iceberry
 * @date 2022-1-2
 * @version 1.0
 * map
 * */
package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	basicUse()
	hasKey()
	traversalMap()
	deleteMap()
	traversalMapOrderly()
	mapValueSlice()
	sliceValueMap()
}

func sliceValueMap() {
	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap)
	fmt.Println("after init:")
	key := "中国"
	value, ok := sliceMap[key]
	if !ok {
		value = make([]string, 0, 2)
	}
	value = append(value, "北京", "上海")
	sliceMap[key] = value
	fmt.Println(sliceMap)
	printSep()
}

func mapValueSlice() {
	var mapSlice = make([]map[string]string, 3)
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	fmt.Println("after init:")
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "王五"
	mapSlice[0]["password"] = "123456"
	mapSlice[0]["address"] = "北京"
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	printSep()
}

func traversalMapOrderly() {
	rand.Seed(time.Now().UnixNano())
	var scoreMap = make(map[string]int, 200)
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100)
		scoreMap[key] = value
	}
	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
	printSep()
}

func deleteMap() {
	scoreMap := make(map[string]int, 8)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	scoreMap["王五"] = 60
	delete(scoreMap, "小明")
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}
	printSep()
}

func traversalMap() {
	scoreMap := make(map[string]int, 8)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	scoreMap["王五"] = 60
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}
	for k := range scoreMap {
		fmt.Println(k)
	}
	printSep()
}

func hasKey() {
	scoreMap := make(map[string]int, 8)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	v, ok := scoreMap["张三"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("查无此人")
	}
	printSep()
}

func basicUse() {
	scoreMap := make(map[string]int, 8)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["小明"])
	fmt.Printf("type of scoreMap:%T\n", scoreMap)
	userInfo := map[string]string{
		"username": "Iceberry",
		"password": "123456",
	}
	fmt.Println(userInfo)
	printSep()
}

func printSep() {
	fmt.Println("----------------------------------------")
}
