package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// listen for incoming connections.
	l, err := net.Listen("tcp", "localhost:7000")
	if err != nil {
		fmt.Println("failed to listen with error:", err.Error())
		os.Exit(1)
	}
	// close the listener when the application closes.
	defer l.Close()

	fmt.Println("listening on " + CONN_HOST + ":" + CONN_PORT)
	for {
		// listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("failed to accept with error:", err.Error())
			os.Exit(1)
		}
		// handle connections in a new goroutine.
		go handleRequest(conn)
	}
}

// handles incoming requests.
func handleRequest(conn net.Conn) {
	// make a buffer to hold incoming data.
	buf := make([]byte, 1024)

	// read the incoming connection into the buffer.
	readLen, err := conn.Read(buf)
	if err != nil {
		fmt.Println("failed to read with error:", err.Error())
	}

	// send a response back to person contacting us.
	conn.Write(buf[:readLen])

	// close the connection when you're done with it.
	conn.Close()
}
