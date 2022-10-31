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

func NewUser(conn net.Conn) *User {
	addr := conn.RemoteAddr().String()
	u := &User{
		addr,
		addr,
		make(chan string),
		conn,
	}
	// 启动监听当前User channel信息
	go u.ListenMessage()
	return u
}

func (u *User) ListenMessage() {
	for msg := range u.C {
		_, _ = u.conn.Write([]byte(msg + "\n"))
	}
}
