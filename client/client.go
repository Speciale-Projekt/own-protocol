package client

import (
	"bufio"
	"fmt"
	"net"
)

func CreateClient(wrong bool) {
	p := make([]byte, 2048)
	conn, err := net.Dial("udp", "localhost:1234")
	if err != nil {
		fmt.Printf("Some error %v", err)
		return
	}
	if !wrong {
		fmt.Fprintf(conn, "Hi")
		_, err = bufio.NewReader(conn).Read(p)
		if err == nil {
			fmt.Printf("%s\n", p)
		} else {
			fmt.Printf("Some error %v\n", err)
		}
	} else {
		fmt.Fprintf(conn, "volapyk")
		_, err = bufio.NewReader(conn).Read(p)
		if err == nil {
			fmt.Printf("%s\n", p)
		} else {
			fmt.Printf("Some error %v\n", err)
		}
	}
	conn.Close()
}
