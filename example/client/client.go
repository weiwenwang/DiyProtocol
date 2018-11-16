package main

import (
	"net"
	"fmt"
	"encoding/binary"
	"bytes"
)

func main() {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", "127.0.0.1", "8080"))
	if err != nil {
		panic(err.Error())
	}
	hw := []string{
		"于IP协议 ，IP IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就" +
			"是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长" +
			"度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据" +
			"包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（" +
			"IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MT" +
			"U决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP" +
			"包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长 ",
		"于IP协议 ，IP IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就" +
			"是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长" +
			"度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据" +
			"包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（" +
			"IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MT" +
			"U决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP" +
			"包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长 ",
		"于IP协议 ，IP IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就" +
			"是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长" +
			"度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据" +
			"包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（" +
			"IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MT" +
			"U决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP" +
			"包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长 ", "于IP协议 ，IP IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就" +
			"是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长" +
			"度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据" +
			"包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（" +
			"IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MT" +
			"U决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP" +
			"包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长 ", "于IP协议 ，IP IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就" +
			"是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长" +
			"度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据" +
			"包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（" +
			"IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MT" +
			"U决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP" +
			"包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长 ", "于IP协议 ，IP IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就" +
			"是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长" +
			"度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据" +
			"包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（" +
			"IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MT" +
			"U决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP" +
			"包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长 ", "于IP协议 ，IP IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就" +
			"是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长" +
			"度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据" +
			"包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（" +
			"IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MT" +
			"U决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP" +
			"包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长 ",
		"于IP协议 ，IP IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就" +
			"是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长" +
			"度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据" +
			"包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（" +
			"IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MT" +
			"U决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP" +
			"包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长 ",
		"于IP协议 ，IP IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就" +
			"是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长" +
			"度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据" +
			"包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（" +
			"IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MT" +
			"U决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP" +
			"包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长 ",
		"于IP协议 ，IP IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就" +
			"是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长" +
			"度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据" +
			"包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（" +
			"IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MT" +
			"U决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP" +
			"包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长 ", "于IP协议 ，IP IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就" +
			"是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长" +
			"度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据" +
			"包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（" +
			"IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MT" +
			"U决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP" +
			"包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长 ", "于IP协议 ，IP IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就" +
			"是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长" +
			"度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据" +
			"包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（" +
			"IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MT" +
			"U决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP" +
			"包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长 ", "于IP协议 ，IP IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就" +
			"是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长" +
			"度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据" +
			"包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（" +
			"IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MT" +
			"U决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP" +
			"包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长 ", "于IP协议 ，IP IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就" +
			"是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长" +
			"度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据" +
			"包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（" +
			"IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MT" +
			"U决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP" +
			"包的大小由MTU决定（IP数据包长度就是于IP协议来说，IP包的大小由MTU决定（IP数据包长 ",
	}
	for _, v := range hw {
		headSize := len(v)
		var headBytes = make([]byte, 2)
		binary.BigEndian.PutUint16(headBytes, uint16(headSize))
		var buffer bytes.Buffer

		buffer.Write([]byte("BEGIN"))
		buffer.Write(headBytes)
		buffer.WriteString(v)
		b3 := buffer.Bytes() //得到了b1+b2的结果
		conn.Write(b3)

		fmt.Println("headSize = ", len(b3))
	}
}
