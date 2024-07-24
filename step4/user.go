package main

import (
	"net"
)

type Usr struct {
	Name string // user name
	Addr string // user remote ip address
	C    chan string
	conn net.Conn
}

// initiate a new user
func NewUsr(conn net.Conn) *Usr {
	userAddr := conn.RemoteAddr().String()
	usr := &Usr{
		Name: userAddr,
		Addr: userAddr,
		C:    make(chan string),
		conn: conn,
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
