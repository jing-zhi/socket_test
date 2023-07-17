package main

import (
	"fmt"
	"net"
)

func main() {
	// 创建一个监听器，监听本地的 8000 端口,listener 是

	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}

	defer listener.Close()

	fmt.Println("Listening on :8000")

	for {
		// 等待客户端连接
		//Client.Cli()
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting:", err.Error())
			continue
		}

		// 处理客户端连接

		go handleConnection(conn)

	}

}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	fmt.Println("Client connected:", conn.RemoteAddr())

	buf := make([]byte, 1024)
	for {
		// 从客户端读取数据
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			return
		}

		// 打印客户端发送的数据
		fmt.Println("Received data:", string(buf[:n]))

		// 将数据回传给客户端
		_, err = conn.Write(buf[:n])
		if err != nil {
			fmt.Println("Error writing:", err.Error())
			return
		}
	}
}
