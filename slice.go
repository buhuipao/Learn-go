package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := a
	fmt.Println(a)
	fmt.Println(b)
	// 当引用的slice的长度大于原slice时，将重新指向一个底层slice
	b = append(b, 5)
	fmt.Println(a)
	fmt.Println(b)
	s := make([]int, 5, 10)
	fmt.Println(s)
	changeSlice(s)
	fmt.Println(s)
}

func changeSlice(s []int) {
	s[3] = 1000
}
