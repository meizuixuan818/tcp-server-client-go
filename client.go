package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	fmt.Println("client start!\n",
		"按下q退出……\n",
		"输入发送的内容后按下回车键发送！\n")
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer conn.Close() //关闭连接
	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, _ := inputReader.ReadString('\n') //读取用户输入
		inputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputInfo) == "Q" { //输入q/Q就退出
			return
		}
		_, err = conn.Write([]byte(inputInfo)) // 发送数据
		if err != nil {
			return
		}
		buf := [512]byte{}
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("recv failed,err:", err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}
