package server

import (
	"fmt"
	"log"
	"net"
)

func ServerInit(addr string) error {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Errorf("%s\n", err.Error())
		return err
	}
	fmt.Printf("server start success....\n")

	for {
		c, errAt := ln.Accept()
		if errAt != nil {
			fmt.Errorf("accept error:%s\n", errAt.Error())
			break
		}
		go handleConn(c)
	}
	return nil
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		var buf = make([]byte, 100)
		log.Println("start to read from conn")
		n, err := c.Read(buf)
		if err != nil {
			log.Println("conn read error:", err)
			return
		}
		log.Printf("read %d bytes, content is %s\n", n, string(buf[:n]))

		var ret = "welcome..."
		n, err = c.Write([]byte(ret))
		if err != nil {
			log.Printf("write err,buf:%s\n", ret)
			return
		}
		log.Printf("write success, len:%d, buf:%s\n", n, ret)

	}
}
