package main

import (
	"fmt"
	"net"
	"os"
	"DiyProtocol/nianbao"
)

func main() {
	addr := fmt.Sprintf("%s:%s", "127.0.0.1", "8080")
	fmt.Println(addr)
	listen, err1 := net.Listen("tcp", addr)
	if err1 != nil {
		fmt.Println("Error: %s", err1.Error())
		os.Exit(1)
	}
	for {
		conn, err2 := listen.Accept()
		if err2 != nil {
			fmt.Println("Error: %s", err2.Error())
			os.Exit(1)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	msg := make(chan string, 10)             // 这里设置消息channel可以容纳10个消息
	buffer1 := nianbao.NewBuffer(conn, 1024) // 缓存区设置1024字节， 如果单个消息大于这个值就不能接受了
	for {
		select {
		case msg1 := <-msg: // 从msg chan里面取数据
			doMsg(conn, msg1)
			// 读到数据放到msg chan里面, 如果msg里没有消息了，那么会阻塞在ThomasRead这个函数里等消息
			// 当读到数据了，并分析这次读到的数据里有几个消息, 不足一个消息就继续读，大于一个消息了，处理掉这个消息
		default:
			buffer1.Read(msg)
		}
	}
}

func doMsg(conn net.Conn, str string) {
	// conn.Write() 这里可能要发送点什么到对等方
	fmt.Println("length:", len(str))
}
