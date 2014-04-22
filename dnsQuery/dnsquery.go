package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		addr, err := net.LookupIP(os.Args[i])
		if err != nil {
			fmt.Printf("Parsing %s error, and the error is %s\n", os.Args[i], err)
		} else {
			fmt.Printf("The IP addresses of %s are:\n", os.Args[i])
			for _, arr := range addr {
				fmt.Printf("	%s\n", arr)
			}
		}
	}
}