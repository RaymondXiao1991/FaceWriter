package common

import (
	"reflect"
)

func SliceEqual(a, b interface{}) bool {
	voa, vob := reflect.ValueOf(a), reflect.ValueOf(b)
	if voa.Kind() != reflect.Slice || vob.Kind() != reflect.Slice {
		panic("Param $1 and $2 must be a slice")
	}

	if voa.Len() != vob.Len() {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	for i := 0; i < voa.Len(); i++ {
		if voa.Index(i).Interface() != vob.Index(i).Interface() {
			return false
		}
	}
	return true
}

type BySortIndex []int64

func (b BySortIndex) Len() int { return len(b) }

func (b BySortIndex) Swap(i, j int) { b[i], b[j] = b[j], b[i] }

func (b BySortIndex) Less(i, j int) bool {
	return b[i] < b[j]
}

type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less reports whether the element with
	// index i should sort before the element with index j.
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}

func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

// 实现类似python的迭代器
func Generator() chan int {
	c := make(chan int)
	n := 0
	go func() {
		for {
			c <- n
			n = n + 1
		}
	}()

	return c
}

func Abs(n int64) int64 {
	y := n >> 63
	return (n ^ y) - y
}

// Reverse 反转字符串
func Reverse(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}
