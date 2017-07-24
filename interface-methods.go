package main

import (
	"fmt"
)

type List []int

func (l List) Len() int {
	return len(l)
}

func (l *List) Append(i int) {
	*l = append(*l, i)
}

type Appender interface {
	Append(int)
}

type Lenner interface {
	Len() int
}

func AppendItem(a Appender, start, end int) {
	for i := start; i <= end; i++ {
		a.Append(i)
	}
}

func GetLen(l Lenner) int {
	return l.Len()
}

func main() {
	// 值传递
	var l Appender
	list := new(List)
	l = list
	AppendItem(l, 1, 10)
	var L Lenner
	L = list
	fmt.Printf("Length of %d is %d", L, L.Len())

	// 传指针的接口支持传值，反之不行
	p_list := new(List)
	l = p_list
	AppendItem(l, 10, 100)
	L = list
	fmt.Printf("Length of %d is %d", L, GetLen(L))
}
