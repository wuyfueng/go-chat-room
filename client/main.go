package main

import (
	"flag"
	"log"
	"net"
	"time"

	"github.com/wuyfueng/go-chat-room/common"
	"github.com/wuyfueng/tools"
)

var (
	msg          = flag.String("msg", "hello", "发送的消息")
	sendDuration = flag.Int("sendDuration", 3, "发送间隔(秒)")
)

func main() {
	flag.Parse()

	// 建立连接
	conn, err := net.DialTCP("tcp", nil, common.TcpAddr)
	if err != nil {
		log.Fatalln("ListenTCP err", err)
	}

	for range time.Tick(time.Second * time.Duration(*sendDuration)) {
		if sendAndRecv(conn) {
			break
		}
	}

	tools.WaitExit()
}

// 收发消息
func sendAndRecv(conn *net.TCPConn) (isClose bool) {
	// 发送消息
	_, err := conn.Write([]byte(*msg))
	if err != nil {
		log.Printf("Write RemoteAddr: %v err: %v", conn.RemoteAddr().String(), err)
		return true
	}
	log.Println("发送消息", *msg)

	// 接收消息
	var buf [128]byte
	_, err = conn.Read(buf[:])
	if err != nil {
		log.Printf("Read RemoteAddr: %v err: %v", conn.RemoteAddr().String(), err)
		return true
	}
	log.Printf("来自远端: %s, 消息: %s, ", conn.RemoteAddr().String(), string(buf[:]))

	return false
}
