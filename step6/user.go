package main

import (
	"net"
	"strings"
)

// step4: add server struct to server
type Usr struct {
	Name string // user name
	Addr string // user remote ip address
	C    chan string
	conn net.Conn

	server *Server // connect user to server
}

// initiate a new user
func NewUsr(conn net.Conn, srv *Server) *Usr {
	userAddr := conn.RemoteAddr().String()
	usr := &Usr{
		Name: userAddr,
		Addr: userAddr,
		C:    make(chan string),
		conn: conn,
		// step 4: connect user to server
		server: srv,
	}

	//  start listening to any data sent to the user's C data string stream
	// when there is data given to the user's C channel, then it will write to the connection
	go usr.ListenAndWriteToConn()

	return usr
}

// Listens to any incoming message coming to the C channel
// writes out the incoming message to the established connection
func (this *Usr) ListenAndWriteToConn() {
	for {
		msg := <-this.C
		_, err := this.conn.Write([]byte(msg + "\n"))
		if err != nil {
			return
		}
	}
}

// deal with user coming online
func (this *Usr) Online() {
	this.server.mapLock.Lock()
	this.server.OnlineMap[this.Name] = this
	this.server.mapLock.Unlock()

	this.server.InitiateBroadcastWithMsg(this, "comes online")
}

// deal with user coming offline
func (this *Usr) Offline() {
	this.server.mapLock.Lock()
	delete(this.server.OnlineMap, this.Name)
	this.server.mapLock.Unlock()

	this.server.InitiateBroadcastWithMsg(this, "comes offline")
}

// deal with user type in messages to be broadcast to the group
func (this *Usr) GroupMessage(msg string) {
	if msg == "who" {
		this.server.mapLock.Lock()
		for _, u := range this.server.OnlineMap {
			this.conn.Write([]byte(u.Name + " " + u.Addr + " is online\n"))
		}
		this.server.mapLock.Unlock()
	} else if len(msg) > 7 && msg[:7] == "rename|" {
		newName := strings.Split(msg, "|")[1]

		this.server.mapLock.Lock()
		_, ok := this.server.OnlineMap[newName]
		if !ok {
			delete(this.server.OnlineMap, this.Name)
			this.server.OnlineMap[newName] = this
			this.Name = newName
			this.conn.Write([]byte("name changed\n"))
		} else {
			this.conn.Write([]byte("the name already exists\n"))
		}
		this.server.mapLock.Unlock()
	}
	this.server.InitiateBroadcastWithMsg(this, msg)
}