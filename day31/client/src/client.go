/**
 * Created by visual studio code
 * @author Iceberry
 * @date 2022-1-8
 * @version 1.0
 * tcp客户端
 * */
package main

import (
	"fmt"
	"net"
)

func main() {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("connected failed!,err:", err)
		return
	}
	defer socket.Close()
	sendData := []byte("Hello server")
	_, err = socket.Write(sendData)
	if err != nil {
		fmt.Println("send failed,err:", err)
		return
	}
	data := make([]byte, 4096)
	n, remoteAddr, err := socket.ReadFromUDP(data)
	if err != nil {
		fmt.Println("receive data failed!,err:", err)
		return
	}
	fmt.Printf("receive: %v addr: %v count: %v\n", string(data[:n]), remoteAddr, n)
}
