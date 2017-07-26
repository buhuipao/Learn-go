package main

import "fmt"

var i = 5
var str = "ABC "

type Person struct {
	Name string
	Age  int
}

type Any interface{}

func main() {
	var any Any
	any = 19
	fmt.Printf("My age is %v\n", any)
	any = "chenhua"
	fmt.Printf("My name is %v\n", any)
	person := new(Person)
	person.Name = "qiusha"
	person.Age = 18
	any = person
	fmt.Printf("My info is: %v\n", any)
	switch t := any.(type) {
	case int:
		fmt.Printf("Type int %T\n", t)
	case string:
		fmt.Printf("Type string %T\n", t)
	case bool:
		fmt.Printf("Type string %T\n", t)
	case *Person:
		fmt.Printf("Type pointer of Person %T\n", t)
	default:
		fmt.Println("Unknown Type %T", t)
	}

	// 结合匿名函数使用
	TypeSwitch()

	//构建一个二叉树
	BuildTree()
}

type specialString string

var whatIsThis specialString = "hello"

func TypeSwitch() {
	testFunc := func(any interface{}) {
		switch v := any.(type) {
		case bool:
			fmt.Printf("any %v is a bool type\n", v)
		case int:
			fmt.Printf("any %v is an int type\n", v)
		case float32:
			fmt.Printf("any %v is a float32 type\n", v)
		case string:
			fmt.Printf("any %v is a string type\n", v)
		case specialString:
			fmt.Printf("any %v is a special String!\n", v)
		default:
			fmt.Println("unknown type!\n")
		}
	}
	testFunc(whatIsThis)
}

type TreeNode struct {
	left  *TreeNode
	val   interface{}
	right *TreeNode
}

func InitNode(l *TreeNode, val interface{}, r *TreeNode) *TreeNode {
	return &TreeNode{l, val, r}
}

func BuildTree() *TreeNode {
	var val interface{}
	val = "root"
	root := InitNode(nil, val, nil)
	val = "left"
	left := InitNode(nil, val, nil)
	val = "right"
	right := InitNode(nil, val, nil)
	root.left = left
	root.right = right
	fmt.Printf("root's type is %T", root)
	return root
}
