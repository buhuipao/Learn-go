package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "0.0.0.0:5000")
	if err != nil {
		fmt.Println("Connect 0.0.0.0:5000 failed")
		return
	}

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please input your name: ")
	// 会把\n读进去
	clientName, err := inputReader.ReadString('\n')
	trimmedClient := strings.Trim(clientName, "\n")
	for {
		fmt.Println("What to send to the server? Input Q to quit!")
		input, _ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\n")
		if trimmedInput == "Q" {
			return
		}
		_, err = conn.Write([]byte(trimmedClient + " says: " + trimmedInput))
	}

}
