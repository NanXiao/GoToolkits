package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {
	interval := 0
	count := 1
	var err error = nil

	if (len(os.Args) != 3) && (len(os.Args) != 5) {
		fmt.Printf("Usage: tcp_keep_alive ip_addr port [interval(seconds, default:0)] [count(run times, default:1)]\n")
		os.Exit(1)
	}

	if len(os.Args) == 5 {
		interval, err = strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Printf("Convert interval error, the error is %s\n", err)
		}
		count, err = strconv.Atoi(os.Args[4])
		if err != nil {
			fmt.Printf("Convert count error, the error is %s\n", err)
		}
	}

	for count > 0 {
		conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", os.Args[1], os.Args[2]))
		if err != nil {
			fmt.Println(err)
		} else {
			conn.Close()
		}
		count--
		time.Sleep(time.Duration(interval) * time.Second)
	}
}
