package main

import (
	"os"
	"os/signal"
	"serverclientexample/server"
	"syscall"
)

func main() {
	go server.CreateServer()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

}
