package main

import (
	"log"
	"net"
)

/*
run in shell:
telnet localhost 12345
*/
func main() {
	listener, _ := net.Listen("tcp", ":12345")
	for {
		conn, _ := listener.Accept()
		// 处理客户端连接。
		go ClientHandler(conn)
	}
}

func ClientHandler(c net.Conn) {
	defer func() {
		if v := recover(); v != nil {
			log.Println("捕获了一个恐慌：", v)
			log.Println("防止了程序崩溃")
		}
		c.Close()
	}()
	panic("未知错误")
	// fmt.Println("正常")
}
