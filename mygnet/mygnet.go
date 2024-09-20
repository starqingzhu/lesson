package mygnet

import (
	"fmt"
	"github.com/panjf2000/gnet"
)

type echoServer struct {
	*gnet.EventServer
}

func (es *echoServer) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	out = frame
	return
}

func Init() {
	echo := new(echoServer)
	fmt.Println("server init begin.....")
	gnet.Serve(echo, "tcp://:9000")
}