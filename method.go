package main

import (
	"fmt"
)

type A struct {
	Name string
	Age  int
}

type MyInt int

func main() {
	a := &A{Name: "chenhua", Age: 90}
	a.Print()
	fmt.Println(a)
	var t MyInt
	t.Increase(100)
	fmt.Println(t)
}

func (a *A) Print() {
	a.Name = "qiusha"
	a.Age += 10
	fmt.Println("AA")
}

func (t *MyInt) Increase(num int) {
	fmt.Println(t)
	*t += MyInt(num)
}
