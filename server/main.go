package server

import (
	"fmt"
	"github.com/wuyfueng/tools"
	"log"
	"math/rand"
	"net"

	"github.com/wuyfueng/go-chat-room/common"
)

// Start 运行服务端
func Start() {
	// 监听端口
	l, err := net.ListenTCP("tcp", common.TcpAddr)
	if err != nil {
		log.Fatalln("ListenTCP err", err)
	}
	log.Println("tcp监听成功! 地址:", l.Addr().String())

	// 接收连接
	accept(l)
}

// 接收连接
func accept(l *net.TCPListener) {
	// 异步
	f := func() {
		// 有新连接就进行处理, 没有新连接就阻塞l.AcceptTCP()
		for {
			conn, err := l.AcceptTCP()
			if err != nil {
				log.Fatalln("AcceptTCP err", err)
			}
			log.Printf("有新连接! 本地地址: %s, 远端地址: %s", conn.LocalAddr().String(), conn.RemoteAddr().String())

			// 异步处理单条链接
			tools.Async(func() { connDeal(conn) })
		}
	}
	tools.Async(f)
}

// 处理单个连接
func connDeal(conn *net.TCPConn) {
	for {
		// 创建缓存区, 并读取接收到的数据
		var buf [128]byte
		_, err := conn.Read(buf[:])
		if err != nil {
			log.Printf("Read RemoteAddr: %v err: %v", conn.RemoteAddr().String(), err)
			break
		}

		// 向客户端发送消息
		sendMsg := []byte(fmt.Sprintf("幸运数字: %d", rand.Intn(10)))
		_, err = conn.Write(sendMsg)
		if err != nil {
			log.Printf("Write RemoteAddr: %v err: %v", conn.RemoteAddr().String(), err)
			break
		}

		log.Printf("%s发来消息: %s, 回应消息: %s", conn.RemoteAddr().String(), string(buf[:]), sendMsg)
	}
}
