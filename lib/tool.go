package lib

// func Load(times int) {
// 	for i := 0; i < times; i++ {
// 		conn, err := net.Dial("tcp", url)
// 		if err != nil {
// 			panic(err)
// 		}
// 		defer conn.Close()

// 		message := "hihi\n"
// 		_, err = conn.Write([]byte(message))
// 		if err != nil {
// 			panic(err)
// 		}

// 		buff := make([]byte, 1024)
// 		_, err = conn.Read(buff)
// 		if err != nil {
// 			panic(err)
// 		}
// 		fmt.Println(string(buff))
// 	}
// }
