package main

import (
	"fmt"
	"net"
	"os"
	"regexp"
)

const (
	url = "localhost:8890"
)

func Echo(conn net.Conn) {
	defer conn.Close()
	buff := make([]byte, 1024)
	_, err := conn.Read(buff)
	if err != nil {
		fmt.Println("conn.Read failed")
		os.Exit(1)
	}
	remoteAddr := conn.RemoteAddr().String()
	fmt.Println(remoteAddr, ": ", string(buff))

	_, err = conn.Write(buff)
	if err != nil {
		panic(err)
	}

}

func isUserExit(message string) bool {
	return regexp.MustCompile(`exit()`).MatchString(message)
}

func HandleConn(conn net.Conn) {
	buff := make([]byte, 1024)
	for {
		_, err := conn.Read(buff)
		if err != nil {
			// fmt.Println("loop")
			continue
		}
		if isUserExit(string(buff)) {
			fmt.Println("user exit")
			conn.Close()
			break
		}

		remoteAddr := conn.RemoteAddr().String()
		fmt.Println(remoteAddr, ": ", string(buff))

		_, err = conn.Write(buff)
		if err != nil {
			panic(err)
		}

	}
}

type User struct {
	Host string
	Port string
	Name string
}

func main() {

	lis, err := net.Listen("tcp", url)
	if err != nil {
		panic(err)
	}
	fmt.Println("multithreaded echo server")
	fmt.Println("lis", url)

	for {
		conn, err := lis.Accept()
		if err != nil {
			fmt.Println("lis.Accent failed")
			os.Exit(1)
		}

		go HandleConn(conn)
		// Echo(conn)

	}

}
