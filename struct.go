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
	// 尝试改变结构体之后, 得出结论：结构体是值传递
	changeStruct(b)
	fmt.Println(b)
	// 一般直接获取地址即可, 而且直接访问不需要反向取址
	test := &A{name: "jiefu", B: B{age: 19}}
	test.age = 1000
	fmt.Println(test)
	change(test)
	fmt.Println(test)
}

func changeStruct(person B) {
	person.name = "jiefu"
}

func change(test *A) {
	test.name = "changed"
}
