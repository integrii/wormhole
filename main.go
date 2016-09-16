package main

import (
	"io"
	"log"
	"net"
)

func main() {
	// Listen on TCP port 2000 on all interfaces.
	l, err := net.Listen("tcp", ":443")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		// Wait for a connection.
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// Handle the connection in a new goroutine.
		// The loop then returns to accepting, so that
		// multiple connections may be served concurrently.
		go func(c net.Conn) {
			// connect to the destination tcp port
			destConn, err := net.Dial("tcp", "localhost:22")
			if err != nil {
				log.Fatal("Error connecting to destination port")
			}
			io.Copy(c, destConn)
			io.Copy(destConn, c)
			// Shut down the connection.
			c.Close()
		}(conn)
	}
}
