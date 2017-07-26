package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Address   []*Address
	Remark    string
}

// 分别写到字符串和文件
func main() {
	// 定义个人住宅和公司地点
	wa := &Address{"work", "beijing", "china"}
	pa := &Address{"private", "changsha", "hunan"}
	vc := VCard{"chen", "hua", []*Address{wa, pa}, "none"}
	// fmt.Printf("My info is: %v", vc)

	// 虽然返回的是[]byte类型，但是实现了String()函数，其实打印时返回的是string([]byte)
	js, _ := json.Marshal(vc)
	fmt.Printf("JSON format: %s", js)

	// 以666的权限模式创建或者写一个文件
	file, _ := os.OpenFile("vcard.json", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	enc := json.NewEncoder(file)
	err := enc.Encode(vc)
	if err != nil {
		log.Println("Error in encoding json to a file")
	}
}
