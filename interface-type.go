package main

import "fmt"
import "math"

type Square struct {
	side float32
}

type Circle struct {
	radius float32
}

type Shaper interface {
	Area() float32
}

type Stringer interface {
	String() string
}
type V struct {
	Name string
}

func main() {
	var areaIntf Shaper
	sq1 := &Square{}
	sq1.side = 5
	areaIntf = sq1

	// 检测类型
	if _, ok := areaIntf.(*Square); ok {
		fmt.Println("Find a square...")
		fmt.Println("Square's area is: ", sq1.Area())
	} else if _, ok := areaIntf.(*Circle); ok {
		fmt.Println("Circle's area is: ", sq1.Area())
		fmt.Println("Find a circle...")
	} else {
		fmt.Println("areaItf does not contain a variable of type Circle")
	}

	// type-switch 检测类型
	classifier(12, 111.2, "aaa", nil)

	// 检测某个值是否实现了某个接口, 本例是个空接口
	var v interface{}
	// var v Stringer
	vv := &V{}
	v = vv
	if _, ok := v.(Stringer); ok {
		fmt.Printf("%d has implate Stringer interface", v)
	} else {
		fmt.Printf("%d hasn't implate Stringer interface", v)
	}
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (c *Circle) Area() float32 {
	return math.Pi * c.radius * c.radius
}

func (v *V) String() string {
	fmt.Println("Hello World!")
	return "Hello World"
}

func classifier(items ...interface{}) {
	for i, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("Param #%d is a bool\n", i)
		case string:
			fmt.Printf("Param #%d is a string\n", i)
		case float64:
			fmt.Printf("Param #%d is a float64\n", i)
		case int, int64:
			fmt.Printf("Param #%d is a int\n", i)
		case nil:
			fmt.Printf("Param #%d is a nil\n", i)
		default:
			fmt.Printf("Param #%d is unknown\n", i)
		}
	}
}
