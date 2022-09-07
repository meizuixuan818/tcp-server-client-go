package main

import (
	"bufio"
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close() //关闭连接
	fmt.Println("connection success!")
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:]) //读取数据
		if err != nil {
			fmt.Println("read from client failed,error:", err)
			break
		} else {
			fmt.Println("success read!")
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client端发送的数据：", recvStr)
		conn.Write([]byte(recvStr)) //发送数据
	}
}

func main() {
	// listen, err := net.Listen("tcp", "127.0.0.1:20000")
	listen, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		fmt.Println("Listen failed,error:", err)
		return
	} else {
		fmt.Println("success listen!")
	}
	fmt.Println("waiting for client")
	for {
		conn, err := listen.Accept() //建立连接
		if err != nil {
			fmt.Println("accept failed,error:", err)
			continue
		} else {
			fmt.Println("success accept!")
		}
		go process(conn) //启动一个goroutine处理连接
	}
}
