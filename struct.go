package main

import (
	"fmt"
)

type A struct {
	name string
	B
}

type B struct {
	name string
	age  int
}

func main() {
	// 使用嵌入结构时，默认使用一级字段，再使用二级字段(应避免冲突)
	a := A{name: "chenhua", B: B{name: "qiusha"}}
	b := B{
		name: "chenhua",
		age:  110,
	}
	fmt.Println(a.name, a.B.name)
	fmt.Println(b)
	// 尝试改变结构体之后, 得出结论：结构体传递的值
	changeStruct(b)
	fmt.Println(b)
}

func changeStruct(person B) {
	person.name = "jiefu"
}
