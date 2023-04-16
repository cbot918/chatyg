package main

import (
	"fmt"
	"os"

	"github.com/cbot918/chatyg/lib"
)

const (
	url = "localhost:8890"
)

func main() {
	var userName string
	args := os.Args
	if len(args) == 1 || len(args) != 2 {
		fmt.Println("command: go run . name ")
		return
	}
	userName = args[1]

	fmt.Println("user: ", userName)
	fmt.Println("welcome to tcp client, type exit() to quit")

	chatCli := lib.InitChatCli(url, userName)
	chatCli.Message()
}
