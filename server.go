package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:5000")
	if err != nil {
		fmt.Println("Error listen")
		return
	}
	fmt.Println("Server started, listen on 0.0.0.0:5000")

	// 监听并接受客户端的请求
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error in Accepting connect", err.Error())
			return
		}
		go ProcessConn(conn)
	}
}

func ProcessConn(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		// 返回都到数据的长度
		len, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading conn", err.Error())
			return
		}
		fmt.Printf("Received data: %v\n", string(buf[:len]))
	}
}
