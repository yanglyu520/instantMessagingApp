package main

import (
	"net"
)

type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn
}

// initiate a new user
func NewUser(conn net.Conn) *User {
	userAddr := conn.RemoteAddr().String()
	usr := &User{
		Name: userAddr,
		Addr: userAddr,
		C:    make(chan string),
		conn: conn,
	}

	//  start listening to user sending messages
	go usr.ListenMessage()

	return usr
}

func (this *User) ListenMessage() {
	for {
		msg := <-this.C
		this.conn.Write([]byte(msg + "\n"))
	}
}
