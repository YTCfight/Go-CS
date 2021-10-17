package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	conn, _ := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 8000,
	})
	defer conn.Close()

	content := make([]byte, 1024)
	for {
		n, addr, _ := conn.ReadFromUDP(content)
		fmt.Print("服务端收到的内容：", string(content[:n]))
		c := strings.ToUpper(string(content[:n]))
		conn.WriteToUDP([]byte(c), addr)
	}
}
