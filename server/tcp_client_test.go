package server

import "testing"

func TestTcpClient(t *testing.T) {
	ln, err := InitClient("", 8081)
	if err != nil {
		t.Fatal(err.Error())
	}
	defer ln.Close()

	var buf = "hello world"

	var n int
	n, err = ln.Write([]byte(buf))
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Logf("write to server, %s\n", buf)

	var ret = make([]byte, 100)
	n, err = ln.Read(ret)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Logf("read success, len:%d, buf:%s\n", n, ret[:n])
}
