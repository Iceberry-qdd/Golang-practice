/**
 * Created by visual studio code
 * @author Iceberry
 * @date 2022-1-5
 * @version 1.0
 * 单元测试
 * */
package test

import (
	fib "day22/src/Fib"
	"testing"
)

func benchmarkFib(b *testing.B, n int) {
	//time.Sleep(5 * time.Second)
	//b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fib.Fib(n)
	}
}

func BenchmarkFib1(b *testing.B) { benchmarkFib(b, 1) }

func BenchmarkFib2(b *testing.B) { benchmarkFib(b, 2) }

func BenchmarkFib3(b *testing.B) { benchmarkFib(b, 3) }

func BenchmarkFib10(b *testing.B) { benchmarkFib(b, 10) }

func BenchmarkFib20(b *testing.B) { benchmarkFib(b, 20) }

func BenchmarkFib40(b *testing.B) { benchmarkFib(b, 40) }
