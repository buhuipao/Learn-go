package main

import (
	"fmt"
	"sort"
)

func main() {
	map1()
	map2()
	map3()
	test := map[int]string{1: "1"}
	changeMap(test)
	fmt.Println(test)
}

func map1() {
	var m map[int]map[int]string
	m = make(map[int]map[int]string)
	// 多级map嵌套时，必须每级都需要初始化
	// m[1][1] = "OK" 将会出错
	a, ok := m[1][1]
	if !ok {
		m[1] = make(map[int]string)
	}
	m[1][1] = "OK"
	a = m[1][1]
	fmt.Println(a)
}

func map2() {
	// 初始化一个slice，元素类型为map
	sm := make([]map[int]string, 5)
	for i, v := range sm {
		// rang取到的是值而不是引用，所以无法改变和赋值
		v = make(map[int]string, 1)
		v[1] = "100"
		sm[i] = make(map[int]string, 1)
		sm[i][2] = "200"
		fmt.Println(sm[i], v)
	}
	fmt.Println(sm)
}

func map3() {
	m := map[int]string{1: "a", 3: "c", 2: "b", 4: "d"}
	s := make([]int, len(m))
	i := 0
	for k, _ := range m {
		s[i] = k
		i++
	}
	sort.Ints(s)
	fmt.Println(s)
}

func changeMap(test map[int]string) {
	test[1] = "111"
}
