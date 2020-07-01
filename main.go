package main

import (
	"flag"
	"io"
	"log"
	"net"
	"time"
)

var from *string
var to *string

func init() {
	from = flag.String("from", "0.0.0.0:443", "The address and port that wormhole should listen on.  Connections enter here.")
	to = flag.String("to", "127.0.0.1:80", "Specifies the address and port that wormhole should redirect TCP connections to.  Connections exit here.")
	flag.Parse()
}

func main() {
	log.Println("Starting up.  Forwarding connections from", *from, "to", *to)

	// Listen on the specified TCP port on all interfaces.
	l, err := net.Listen("tcp", *from)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		// Wait for a connection.
		c, err := l.Accept()
		if err != nil {
			log.Println("Error establshing incoming connection:", err)
			continue
		}
		log.Println("Client connected from", c.RemoteAddr())

		// handle the connection in a goroutine
		go wormhole(c)
	}
}

// wormhole opens a wormhole from the client connection
// to user the specified destination
func wormhole(c net.Conn) {
	defer c.Close()
	log.Println("Opening wormhole from", c.RemoteAddr())
	start := time.Now()

	// connect to the destination tcp port
	destConn, err := net.Dial("tcp", *to)
	if err != nil {
		log.Println("Error connecting to destination port:", err)
		return
	}
	defer destConn.Close()
	log.Println("Wormhole open from", c.RemoteAddr())

	go func() { io.Copy(c, destConn) }()
	io.Copy(destConn, c)

	end := time.Now()
	duration := start.Sub(end)
	log.Println("Closing wormhole from", c.RemoteAddr(), "after", duration)
}
