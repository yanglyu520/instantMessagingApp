package main

import (
	"fmt"
	"net"
	"sync"
)

// a tcp connection is between two processes from client to server
// for the server side, we need to know which ip and port to bind and then start listening for any incoming connection request
// therefore, a server needs to have IP address and port for its class definition
type Server struct {
	Ip   string
	Port int
	// online users map
	OnlineMap map[string]*User
	mapLock   sync.RWMutex

	// broadcasting messaging channel
	Message chan string
}

// initiate a new server
func NewServer(ip string, port int) *Server {
	return &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
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

	go this.ListenMessager()
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
	// 1. the user is online put the user into the online map
	// create the new user
	user := NewUser(conn)
	// add the user to the map
	this.mapLock.Lock()
	this.OnlineMap[user.Name] = user
	this.mapLock.Unlock()
	// 2. broadcast that the user comes online
	this.Broadcast(user, " user is online")
}

func (this *Server) Broadcast(user *User, msg string) {
	sendMsg := fmt.Sprintf("user %s", user.Name)
	this.Message <- sendMsg
}

func (this *Server) ListenMessager() {
	for {
		msg := <-this.Message

		// broadcast the message to all users online
		this.mapLock.Lock()
		for _, usr := range this.OnlineMap {
			usr.C <- msg
		}
		this.mapLock.Unlock()
	}
}
