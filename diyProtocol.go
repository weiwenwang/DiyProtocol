package DiyProtocol

import (
	"fmt"
	"encoding/binary"
	"io"
	"errors"
	"os"
	"time"
)

const HEAD_SIZE = 2

type Buffer struct {
	reader io.Reader
	header string
	buf    []byte
	start  int
	end    int
}

func NewBuffer(reader io.Reader, header string, len int) *Buffer {
	buf := make([]byte, len)
	return &Buffer{reader, header, buf, 0, 0}
}

// grow 将有用的字节前移
func (b *Buffer) grow() {
	if b.start == 0 {
		return
	}
	copy(b.buf, b.buf[b.start:b.end])
	b.end -= b.start
	b.start = 0
}

func (b *Buffer) len() int {
	return b.end - b.start
}

//返回n个字节，而不产生移位
func (b *Buffer) seek(n int) ([]byte, error) {
	if b.end-b.start >= n {
		buf := b.buf[b.start : b.start+n]
		return buf, nil
	}
	return nil, errors.New("not enough")
}

//舍弃offset个字段，读取n个字段
func (b *Buffer) read(offset, n int) ([]byte) {
	b.start += offset
	buf := b.buf[b.start : b.start+n]
	b.start += n
	return buf
}

//从reader里面读取数据，如果reader阻塞，会发生阻塞
func (b *Buffer) readFromReader() (int, error) {
	n, err := b.reader.Read(b.buf[b.end:])
	//fmt.Println("本次读取数据的长度:", n)
	if (err != nil) {
		if (err.Error() == "EOF") { // 发送端断开了
			fmt.Println("对等方关闭了")
			os.Exit(1)
		}
		fmt.Println(err.Error())
		panic(err.Error())
		return n, err
	}
	time.Sleep(1 * time.Second) // 便于观察这里sleep了一下
	b.end += n
	return n, nil
}

func (buffer *Buffer) Read(msg chan string) (error) {
	for {
		buffer.grow()                     // 移动数据
		_, err := buffer.readFromReader() // 读数据拼接到定额缓存后面
		if err != nil {
			fmt.Println(err)
			return err
		}
		// 检查定额缓存里面的数据有几个消息(可能不到1个，可能连一个消息头都不够，可能有几个完整消息+一个消息的部分)
		isBreak := buffer.checkMsg(msg)
		// 只要读到有完整的消息, isBreak就为true, 跳出去处理
		if (isBreak) {
			return nil
		}
	}
}

func (buffer *Buffer) checkMsg(msg chan string) (bool) {
	var isBreak bool
	HEADER_LENG := HEAD_SIZE + len(buffer.header)
	headBuf, err1 := buffer.seek(HEADER_LENG)
	if err1 != nil { // 一个消息头都不够， 跳出去继续读吧
		return false
	}
	if (string(headBuf[:len(buffer.header)]) == buffer.header) { // 判断消息头正确性

	} else {

	}
	contentSize := int(binary.BigEndian.Uint16(headBuf[len(buffer.header):]))
	if (buffer.len() >= contentSize-HEADER_LENG) { // 一个消息体也是够的
		contentBuf := buffer.read(HEADER_LENG, contentSize) // 把消息读出来，把start往后移
		msg <- string(contentBuf)
		// 递归，看剩下的还够一个消息不
		isBreak = true
		buffer.checkMsg(msg)
	} else { // 一个消息体不够的， 跳出去继续读吧
		isBreak = false
	}
	return isBreak
}
