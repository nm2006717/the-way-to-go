package main

import (
	"fmt"
	"io"
	"net"
	"sync"
)

type Server struct {
	Ip   string
	Port int

	//在线用户列表
	OnlineMap map[string]*User
	mapLock   sync.RWMutex

	//消息广播的channel
	Message chan string
}

//创建一个server的接口
func NewServer(ip string, port int) *Server {
	var server = new(Server)
	server.Ip = ip
	server.Port = port
	server.OnlineMap = make(map[string]*User)
	server.Message = make(chan string)
	return server
}

func (server *Server) Handler(conn net.Conn) {
	//...当前链接的业务
	// fmt.Println("链接创建成功")

	//创建用户
	var user = NewUser(conn, server)

	//用户上线
	user.Online()

	// 接收客户端发送的消息

	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				user.Offline()
				return
			}
			if err != nil && err != io.EOF {
				fmt.Println("Conn Read err:", err)
				return
			}

			//提取用户的消息(去除'\n')
			msg := string(buf[:n-1])

			//用户针对msg进行消息处理
			user.DoMessage(msg)
		}
	}()

	//当前handle阻塞
	select {}

}

func (server *Server) broadCast(user *User, msg string) {
	var sendMsg = "[" + user.Addr + "]" + user.Name + ":" + msg

	server.Message <- sendMsg

}

//监听Message广播消息channel的goroutine，一旦有消息就发送给全部的在线User
func (server *Server) MessageListener() {
	for {
		msg := <-server.Message

		//将msg发送给全部的在线User
		server.mapLock.Lock()
		for _, user := range server.OnlineMap {
			user.C <- msg
		}
		server.mapLock.Unlock()
	}

}

func (server *Server) Start() {
	// socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", server.Ip, server.Port))
	if err != nil {
		fmt.Println("net.listen err:", err)
		return
	}
	//close listen socket
	defer listener.Close()

	//启动监听Mesage的goroutine
	go server.MessageListener()

	for {
		//accpet
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept err:", err)
			continue
		}

		// do handler
		go server.Handler(conn)
	}

}
