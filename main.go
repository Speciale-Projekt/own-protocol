package main

import (
	"net"
	"os"
	"os/signal"
	"serverclientexample/server"
	"syscall"
)

var (
	addr net.UDPAddr
	ser  *net.UDPConn
)

func init() {
	// Register the signal handler
	addr = net.UDPAddr{
		Port: 1234,
		IP:   net.ParseIP("localhost"),
	}
	ser, _ = net.ListenUDP("udp", &addr)
}

func main() {
	go server.CreateServer(ser)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	ser.Close()

}
