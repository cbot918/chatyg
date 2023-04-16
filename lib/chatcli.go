package lib

import (
	"fmt"
	"net"
	"os"
)

type ChatCli struct {
	UserName string
	Conn     net.Conn
}

func InitChatCli(url string, userName string) *ChatCli {
	c := new(ChatCli)
	var err error
	c.Conn, err = net.Dial("tcp", url)
	if err != nil {
		fmt.Println("net.Dial failed")
		os.Exit(1)
	}

	return c
}

func (c *ChatCli) Message() {

	defer c.Conn.Close()
	for {
		var input string
		fmt.Printf("> ")
		_, err := fmt.Scanln(&input)
		if err != nil {
			panic(err)
		}

		_, err = c.Conn.Write([]byte(input))
		if err != nil {
			panic(err)
		}

		buff := make([]byte, 1024)
		_, err = c.Conn.Read(buff)
		if err != nil {
			if err.Error() == "EOF" {
				fmt.Println("exit program")
				break
			}
			panic(err)
		}
		remoteAddr := c.Conn.RemoteAddr().String()
		fmt.Println(remoteAddr, ": ", string(buff))

	}
}
