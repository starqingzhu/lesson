package mynet

import (
	"net"
)

//获取全部网卡的全部IP
func Ips() (map[string]string, error) {

	ips := make(map[string]string)

	//返回 interface 结构体对象的列表，包含了全部网卡信息
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	//遍历全部网卡
	for _, i := range interfaces {

		// Addrs() 方法返回一个网卡上全部的IP列表
		address, err := i.Addrs()
		if err != nil {
			return nil, err
		}

		//遍历一个网卡上全部的IP列表，组合为一个字符串，放入对应网卡名称的map中
		for _, v := range address {
			ips[i.Name] += v.String() + " "
		}
	}
	return ips, nil
}

func TestGetPublicAddr() {

	var ips = make(map[string]string)
	ips, _ = Ips()

	var ipstring string

	//key是网卡名称，value是网卡IP
	for k, v := range ips {
		ipstring += "testpublicaddr adapter name: " + k + "  IP: " + v + "\n"
	}

}