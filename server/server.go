package server

import (
	"bytes"
	"fmt"
	"net"
)

func sendResponse(conn *net.UDPConn, addr *net.UDPAddr, resp string) {
	_, err := conn.WriteToUDP([]byte(resp), addr)
	if err != nil {
		fmt.Printf("Couldn't send response %v", err)
	}
}

func CreateServer() {
	addr := net.UDPAddr{
		Port: 1234,
		IP:   net.ParseIP("localhost"),
	}
	ser, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Printf("Some error %v\n", err)
		return
	}

	for {
		p := make([]byte, 2048)
		_, remoteaddr, _ := ser.ReadFromUDP(p)
		go HandleShit(ser, remoteaddr, p)
	}
}

func HandleShit(ser *net.UDPConn, remoteaddr *net.UDPAddr, p []byte) {

	p_str := string(bytes.Trim(p, "\x00"))

	// Concert p to string
	switch p_str {
	case "Hi":
		println("hi")
		go sendResponse(ser, remoteaddr, "123")
	default:
		println("Expected 'Hi' got something else ")
		resp := "err"
		go sendResponse(ser, remoteaddr, resp)
	}
}
