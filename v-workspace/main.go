package main

import (
	"fmt"
	"net"
	"os"
)

var arg = os.Args[1]

// TODO: Need to add a clean error message if no arg is passed.

func main() {
	shell(arg)
}

func shell(host string) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		fmt.Println(err)
	} else if conn != nil {
		fmt.Println(&conn)
	}
	// TODO: Need to send some data through the conn object.
}
