package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	socket, _ := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 8000,
	})
	defer socket.Close()

	reply := make([]byte, 1024)
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("请输入")
		msg, _ := reader.ReadString('\n')
		socket.Write([]byte(msg))
		n, _, _ := socket.ReadFromUDP(reply)
		fmt.Print("客户端收到的消息：", string(reply[:n]))
	}
}
