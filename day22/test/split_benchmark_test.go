/**
 * Created by visual studio code
 * @author Iceberry
 * @date 2022-1-5
 * @version 1.0
 * 基准测试
 * */
package test

import (
	split "day22/src/split"
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("即将开始测试……")
	retCode := m.Run()
	fmt.Println("测试完毕，即将退出测试……")
	os.Exit(retCode)
}

func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		split.Split("枯藤老树昏鸦", "老")
	}
}

func BenchmarkSplitParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			split.Split("枯藤老树昏鸦", "老")
		}
	})
}
