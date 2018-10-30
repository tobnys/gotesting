package main

import (
	"fmt"
	"net"
)

func main() {
	srv, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
	}
	for {
		conn, err := srv.Accept()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(conn.LocalAddr())
			fmt.Println(conn.Close())
		}
	}
}
