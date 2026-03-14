package main

// go run main.go -timeout=5s www.google.com 8080
//go run main.go -timeout=5s www.google.com:8080

import (
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	timeout := flag.Duration("timeout", 10 * time.Second, "timeout for connection")
	
	flag.Parse()
	args := flag.Args()
	var host, port string

	if len(args) == 2 {
		host = args[0]
		port = args[1]
	} else if len(args) == 1  {
		var err error
		host, port, err = net.SplitHostPort(args[0])

		if err != nil {
			fmt.Printf("invalid address: %v\n", err)
			os.Exit(1)
		}
	} else {
		fmt.Println("Usage: go-telnet --timeout <duration> <host> <port>")
		os.Exit(1)
	}

	address := net.JoinHostPort(host, port)
	fmt.Println(timeout)
	fmt.Println(address)
}
