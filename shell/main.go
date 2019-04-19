package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:4444") //Send Dial connection to specified ip and port
	log.Fatal(conn, err)

}
