package main

import (
	"fmt"
	"net"
)

// a tcp connection is between two processes from client to server
// for the server side, we need to know which ip and port to bind and then start listening for any incoming connection request
// therefore, a server needs to have IP address and port for its class definition
type Server struct {
	Ip   string
	Port int
}

// initiate a new server
func NewServer(ip string, port int) *Server {
	return &Server{
		Ip:   ip,
		Port: port,
	}
}

// start the server listening and accept connection and send it out to goroutine to handle
func (this *Server) Start() {
	// socket bind and listen, defer closing the listener socket
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net.listener.err: ", err)
		return
	}
	defer listener.Close()
	// accept, blocking function
	for {
		// accept incoming connection request
		// 3 handshakes is happening here where accept func will be sending syn&ack and when receiving client's ack
		// to establish the connection
		// once connection is established, its reading writing to the data stream we built is delegated to a separate goroutine
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("net.accept.err: ", err)
			// ignore error and keep listening for incoming conn request
			continue
		}

		// when accept, we have a http connection
		// we can now handler the http connection in a goroutine
		go this.Handler(conn)
	}
}

func (this *Server) Handler(conn net.Conn) {
	fmt.Println("connection is established", conn)
}
