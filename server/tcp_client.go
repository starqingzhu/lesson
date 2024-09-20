package server

import (
	"fmt"
	"net"
)

func InitClient(ip string, port int) (net.Conn, error) {
	if ip == "" {
		ip = "127.0.0.1"
	}
	addr := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Errorf("dial error:%s, ip:%s, port:%d\n", err.Error(), ip, port)
		return nil, err
	}
	fmt.Printf("connect success ip:%s port:%d...\n", ip, port)

	return conn, nil
}
