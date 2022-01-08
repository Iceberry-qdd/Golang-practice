/**
 * Created by visual studio code
 * @author Iceberry
 * @date 2022-1-7
 * @version 1.0
 * tcp服务端
 * */
package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("listen failed,err:", err)
		return
	}
	fmt.Println("wait for connection……")
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed,err:", err)
			continue
		}

		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Println("read from client failed,err:", err)
			break
		}
		recvStr := string(buf[:])
		fmt.Printf("收到client端发来的长度为%v数据：%v\n", n, recvStr)
		conn.Write([]byte(recvStr))
	}
}
