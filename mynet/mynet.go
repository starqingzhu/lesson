package mynet

import (
	"fmt"
	"github.com/xjh22222228/ip"
)

func GetPublicNetAddrTest() {
	ipv4, err := ip.V4()
	ipv6, err2 := ip.V6()
	if err != nil {
		fmt.Errorf("mynet %s\n", err.Error())
	} else {
		// 98.207.254.136
		fmt.Printf("mynet ipv4 %s\n", ipv4)
	}

	if err2 != nil {
		fmt.Errorf("mynet %s\n", err2.Error())
	} else {
		// 2a00:1450:400f:80d::200e
		fmt.Printf("mynet ipv6 %s\n", ipv6)
	}
}
