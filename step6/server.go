package main

import (
	"fmt"
	"io"
	"net"
	"sync"
)

// step1: build a tcp server that connects between a client and a server process
// a tcp connection is between two processes from client to server
// for the server side, we need to know which ip and port to bind and then start listening for any incoming connection request
// therefore, a server needs to have IP address and port for its class definition

// step2: build a broadcasting function on this server
// once the user is connected to the server connection
// the user's info(like addr) will be stored in the server's map
// then the server will broadcast whoever comes online
type Server struct {
	Ip   string
	Port int
	// online users map
	OnlineMap map[string]*Usr
	mapLock   sync.RWMutex

	// broadcasting messaging channel
	// once a user connects to the server, the server will message to everyone for whoever is online
	Message chan string
}

// initiate a new server
// create an empty map to store users
func NewServer(ip string, port int) *Server {
	return &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*Usr), // store users information
		Message:   make(chan string),
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

	go this.SendMessageToUserChannel()

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
		// step2: the handler connection will create a user upon the connection established
		// the newuser will open another goroutine that listens for any data send to the new user channel
		// step3: read any message from the message channel and start broadcast
		go this.Handler(conn)
	}

}

func (this *Server) Handler(conn net.Conn) {
	// print to stdout
	fmt.Println("connection is established with user: ", conn.RemoteAddr().String())

	// 1. the user is online put the user into the online map
	// create the new user
	// the NewUser method creates a goroutine that listens to any data written to the stream
	user := NewUsr(conn, this)

	user.Online()

	// step3: Reads from any message the client send
	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := conn.Read(buf)
			if err != nil && err != io.EOF {
				fmt.Println("conn.Read err: ", err)
				return
			}
			if n == 0 {
				user.Offline()
				return
			}

			userMsg := string(buf[:n-1])
			user.GroupMessage(userMsg)
		}
	}()

	// block indefinitely
	select {}
}

func (this *Server) InitiateBroadcastWithMsg(user *Usr, msg string) {
	sendMsg := fmt.Sprintf("%s: %s", user.Name, msg)
	this.Message <- sendMsg
}

func (this *Server) SendMessageToUserChannel() {
	for {
		// get message from the message channel
		msg := <-this.Message

		// broadcast the message to all users online
		this.mapLock.Lock()
		for _, u := range this.OnlineMap {
			u.C <- msg
		}
		this.mapLock.Unlock()
	}
}
