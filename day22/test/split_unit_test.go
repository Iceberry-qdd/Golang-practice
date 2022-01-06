/**
 * Created by visual studio code
 * @author Iceberry
 * @date 2022-1-5
 * @version 1.0
 * 单元测试
 * */
package test

import (
	split "day22/src/split"
	"fmt"
	"reflect"
	"testing"
)

func setupTestCase(t *testing.T) func(t *testing.T) {
	t.Log("即将进入测试……")
	return func(t *testing.T) {
		t.Log("测试完毕，即将退出测试……")
	}
}

func setupSubText(t *testing.T) func(t *testing.T) {
	t.Log("即将进入子测试……")
	return func(t *testing.T) {
		t.Log("子测试结束，即将退出子测试……")
	}
}

type test struct {
	input string
	sep   string
	want  []string
}

func TestSplit(t *testing.T) {
	got := split.Split("a:b:c", ":")
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("excepted:%v,got:%v", want, got)
	}
}

func TestMoreSplit(t *testing.T) {
	got := split.Split("abcd", "bc")
	want := []string{"a", "d"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("excepted:%v,got:%v", want, got)
	}
}

func TestSplitGroup(t *testing.T) {
	tests := []test{
		{input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		{input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		{input: "abcd", sep: "bc", want: []string{"a", "d"}},
		{input: "枯藤老树昏鸦", sep: "老", want: []string{"枯藤", "树昏鸦"}},
	}
	for _, tc := range tests {
		got := split.Split(tc.input, tc.sep)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("excepted:%v,got:%v", tc.want, got)
		}
	}
}

func TestSplitGroup2(t *testing.T) {
	tests := map[string]test{
		"simple":      {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		"wrong sep":   {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		"more sep":    {input: "abcd", sep: "bc", want: []string{"a", "d"}},
		"leading sep": {input: "枯藤老树昏鸦", sep: "老", want: []string{"枯藤", "树昏鸦"}},
	}
	for name, tc := range tests {
		got := split.Split(tc.input, tc.sep)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("name:%s excepted:%#v,got:%#v", name, tc.want, got)
		}
	}
}

func TestSplitChild(t *testing.T) {
	tests := map[string]test{
		"simple":      {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		"wrong sep":   {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		"more sep":    {input: "abcd", sep: "bc", want: []string{"a", "d"}},
		"leading sep": {input: "枯藤老树昏鸦", sep: "老", want: []string{"枯藤", "树昏鸦"}},
	}
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			teardownSubTest := setupSubText(t)
			defer teardownSubTest(t)
			got := split.Split(tc.input, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("excepted:%#v, got:%#v", tc.want, got)
			}
		})
	}
}

func ExampleSplit() {
	fmt.Println(split.Split("a:b:c", ":"))
	fmt.Println(split.Split("枯藤老树昏鸦", "老"))
	// Output:
	// [a b c]
	// [枯藤 树昏鸦]
}
