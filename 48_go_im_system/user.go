package main

import (
	"net"
)

type User struct {
	Name   string
	Addr   string
	C      chan string
	conn   net.Conn
	server *Server
}

//创建一个用户的API
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()
	var user = &User{
		Name:   userAddr,
		Addr:   userAddr,
		C:      make(chan string),
		conn:   conn,
		server: server,
	}

	// 启动监听当前user channel消息的goroutine
	go user.MessageListener()

	return user

}

//监听当前User channel的方法，一旦有消息，就直接发送给客户端
func (user *User) MessageListener() {
	for {
		msg := <-user.C
		user.conn.Write([]byte(msg + "\n"))
	}
}

// 用户上线业务
func (user *User) Online() {
	// 用户上线,将用户加入到OnlineMap中
	user.server.mapLock.Lock()
	user.server.OnlineMap[user.Name] = user
	user.server.mapLock.Unlock()
	// 广播当前用户上线消息
	user.server.broadCast(user, "已上线")
}

// 用户下线业务
func (user *User) Offline() {

	// 用户下线,将用户从OnlineMap移除中
	user.server.mapLock.Lock()
	delete(user.server.OnlineMap, user.Name)
	user.server.mapLock.Unlock()
	// 广播当前用户下线消息
	user.server.broadCast(user, "已下线")
}

func (user *User) DoMessage(msg string) {
	if msg == "who" {
		//查看当前在线用户都有谁
		user.server.mapLock.Lock()
		for _, u := range user.server.OnlineMap {
			onlineUserMsg := "[" + u.Addr + "]" + u.Name + ":" + "在线...\n"
			user.sendMsg(onlineUserMsg)
		}
		user.server.mapLock.Unlock()

	} else {
		user.server.broadCast(user, msg)

	}
}

func (user *User) sendMsg(msg string) {
	user.conn.Write([]byte(msg))
}
