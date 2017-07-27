package main

import (
	"errors"
	"fmt"
)

var NotFoundErr error = errors.New("Not found error")

func main() {
	/*
		fmt.Printf("error: %v\n", NotFoundErr)
		defer fmt.Println("After stop program1")
		panic("An Error occurred: stoping... ")
		// 后米这个不会被执行
		defer fmt.Println("After stop program2")
	*/

	// 完整panic-defer<recover>例子
	fmt.Println("Start test...")
	testPanic()
	fmt.Println("End test.")
}

func badCall() {
	panic("bad end")
}
func testPanic() {
	// 在函数最开始定义defer函数，将会在之后出现panic之前运行
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("Logging panic...")
		}
	}()
	badCall()
	// 此后的代码将不会执行
	fmt.Println("After panic, this line won't exec")
}
