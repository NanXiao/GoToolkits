package main

import (
	"fmt"
	"net"
	"os"
)

const (
	maxUDPDatagramLen = 65535
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Usage: udpEchoServer ip_addr port\n")
		os.Exit(1)
	}

	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%s", os.Args[1], os.Args[2]))
	checkError(err)

	conn, err := net.ListenUDP("udp", addr)
	checkError(err)

	buff := make([]byte, maxUDPDatagramLen)

	defer conn.Close()

	for {
		data_len, client, err := conn.ReadFromUDP(buff)
		checkError(err)

		_, err = conn.WriteToUDP(buff[:data_len], client)
		checkError(err)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
