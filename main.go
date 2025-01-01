package main

import (
	"github.com/wuyfueng/go-chat-room/server"
	"github.com/wuyfueng/tools"
)

func main() {
	server.Start()

	tools.WaitExit()
}
