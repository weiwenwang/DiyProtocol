package main

import (
	"net"
	"fmt"
	"encoding/binary"
	"github.com/weiwenwang/DiyProtocol/example"
	"bytes"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", "127.0.0.1", "8080"))
	if err != nil {
		panic(err.Error())
	}
	sendData := getSendData()
	for k, v := range sendData {
		headSize := len(v)
		var headBytes = make([]byte, 2)
		binary.BigEndian.PutUint16(headBytes, uint16(headSize))
		var buffer_client bytes.Buffer

		buffer_client.Write([]byte(example.HEADER))
		buffer_client.Write(headBytes)
		buffer_client.WriteString(v)
		b3 := buffer_client.Bytes() //得到了b1+b2的结果
		_, err := conn.Write(b3)
		if err != nil {
			panic(err)
		}

		fmt.Println("第", k, "个消息", "消息头长:", len(example.HEADER)+len(headBytes), "消息体长:", headSize)
		time.Sleep(500 * time.Millisecond)
	}
}

func getSendData() []string {
	return []string{
		"于IP协议 ，IP IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就" +
			"是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长" +
			"度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据" +
			"包长度就是于IP协议来说，IP包的IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（" +
			"IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MT" +
			"U决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP" +
			"包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长 ",
		"于IP协议 ，IP IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就" +
			"是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长" +
			"度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据" +
			"包长度就是于IP协议来说，包长度就是于IP协议来说，IP包的大小由MTU决定（" +
			"IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MT" +
			"U决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP" +
			"包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长 ",
		"于IP协议 ，IP IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就" +
			"是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长" +
			"度就是于IP协议来说，IP包的大小由MTU决定IP包的大小由MTU决定（IP数据（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据" +
			"包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（" +
			"IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MT" +
			"U决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP" +
			"包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长 ", "于IP协议 ，IP IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就" +
			"是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长" +
			"度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据" +
			"包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（" +
			"IP数据包长度（IP数据包长度就是于IP协议来说，IP包的大小由MT" +
			"U决定（IP数据包长度就是于IP协，IP包的大小由MTU决定（IP议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP" +
			"包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长 ", "于IP协议 ，IP IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就" +
			"是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长" +
			"度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据" +
			"包长度就是于IP协议来说，IP包的大小由MTU就是于IP协议来说，IP包的大小由MTU决定决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（" +
			"IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MT" +
			"U决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP" +
			"包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长 ", "于IP协议 ，IP IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就" +
			"是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长" +
			"度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据" +
			"包长度就是于IP协议来说数据包长度就是于IP协议来说，IP包的大小由MTU决定（" +
			"IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MT" +
			"U决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP" +
			"包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长 ", "于IP协议 ，IP IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就" +
			"是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长" +
			"度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据" +
			"包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（" +
			"IP数据包长度就是于IP协议来说，IP包的大小，IP包的大小由MTU决定（IP由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MT" +
			"U决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP" +
			"包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长 ",
		"于IP协议 ，IP IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就" +
			"是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长" +
			"度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据" +
			"包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（" +
			"IP数据包长度就是于IP协议来说，IP包，IP包的大小由MTU决定（IP的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MT" +
			"U决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP" +
			"包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长 ",
		"于IP协议 ，IP IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就" +
			"是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长" +
			"度就是于IP协议来说，IP包的大小由MT，IP包的大小由MTU决定（IPU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据" +
			"包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（" +
			"IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MT" +
			"U决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP" +
			"包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长 ",
		"nihao包的大小由",
		"于IP协议 ，IP IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就" +
			"是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长" +
			"度就是于IP协议来说，IP包的大，IP包的大小由MTU决定（IP小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据" +
			"包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（" +
			"IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MT" +
			"U决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP" +
			"包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长 ", "于IP协议 ，IP IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就" +
			"是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长" +
			"度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据" +
			"包长度就是于IP协议来说， 定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（" +
			"IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MT" +
			"U决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP" +
			"包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长 ", "于IP协议 ，IP IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就" +
			"是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长" +
			"度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据" +
			"包长度就是于IP协议来说 U决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（" +
			"IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MT" +
			"U决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP" +
			"包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长 ", "于IP协议 ，IP IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就" +
			"是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长" +
			"度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据" +
			"包长度就是于IP协议来说， 大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MT" +
			"U决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP" +
			"包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长 ",
		"nihao包的大小由IP定（IP数据协议来说，I",
		"于IP协议 ，IP IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就" +
			"是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长" +
			"度就是于IP协议来说， 数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据" +
			"包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（" +
			"IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MT" +
			"U决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP" +
			"包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长 ",
	}
}
