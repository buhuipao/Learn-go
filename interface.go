package main

import (
	"bytes"
	"fmt"
)

type Shaper interface {
	Area() float32
	Length() float32
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (sq *Square) Length() float32 {
	sq.side = 10
	return 4 * sq.side
}

type ReadWrit interface {
	Read(b bytes.Buffer) bool
	Write(b bytes.Buffer) bool
}

type Lock interface {
	Lock()
	Unlock()
}

type File interface {
	ReadWrit
	Lock
	Close()
}

func main() {
	// 测试接口
	interface1()
}

func interface1() {
	sq1 := new(Square)
	sq1.side = 5

	var areaIntf Shaper
	areaIntf = sq1
	// shorter,without separate declaration:
	// areaIntf := Shaper(sq1)
	// or even:
	// areaIntf := sq1
	fmt.Printf("The square has area: %f\n", areaIntf.Area())
	fmt.Printf("The square has area: %f\n", areaIntf.Length())
	fmt.Println(sq1.side)
}
