package main

import (
	"fmt"
	"net"
	"os"
)

const (
	url = "localhost:8890"
)

func Load(times int) {
	for i := 0; i < times; i++ {
		conn, err := net.Dial("tcp", url)
		if err != nil {
			panic(err)
		}
		defer conn.Close()

		message := "hihi\n"
		_, err = conn.Write([]byte(message))
		if err != nil {
			panic(err)
		}

		buff := make([]byte, 1024)
		_, err = conn.Read(buff)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(buff))
	}
}

func main() {
	// Load(1)
	var userName string
	args := os.Args
	if len(args) == 1 || len(args) != 2 {
		fmt.Println("command: go run . name ")
		return
	}
	userName = args[1]

	fmt.Println("user: ", userName)
	fmt.Println("welcome to tcp client, type exit() to quit")

	chat := InitChat(url, userName)
	chat.Message()
}
