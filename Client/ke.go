package Client

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func Cli() {
	// 建立到服务器的连接
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	// 发送数据
	fmt.Fprint(conn, "Hello, server!\n")

	// 读取数据
	reader := bufio.NewReader(conn)
	message, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading:", err.Error())
		os.Exit(1)
	}
	fmt.Println("Received message:", message)
}
