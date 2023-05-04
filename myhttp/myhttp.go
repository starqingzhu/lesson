package myhttp

import (
	"fmt"
	"io"
	"io/ioutil"
	"lesson/mysrclib"
	tools "lesson/tools/mytime"
	"net/http"
	"os"
)

func init() {
	fmt.Println("---------welcome myhttp init----------")
}

const (
	URl = "https://www.jd.com/"
)

func TestMyHttp() {

	httpFileName := "./log/http_" + tools.GetFileDateName_Minute()
	httpFileName += ".log"
	httpFile, err := os.OpenFile(httpFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("open file %s failed\n", httpFileName)
		return
	}

	resp, err := http.Get(URl)
	if err != nil {
		mysrclib.MyError.Printf("http get url=%s fail,err=%v", URl, err)
		return
	}
	if len, err := io.Copy(io.Writer(httpFile), resp.Body); err != nil {
		mysrclib.MyError.Printf("get http %s fail,err %v\n", URl, err)
		return
	} else {
		mysrclib.MyInfo.Printf("get http %s success,len %d\n", URl, len)
	}

}

func TestMyAddrHttp() {
	fmt.Println("---------myhttp TestMyAddrHttp----------")
	//addr := "http://10.7.69.9/remote_addr"
	addr := "http://169.254.169.254/latest/meta-data/public-ipv4"
	resp, err := http.Get(addr)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	fmt.Printf("myhttp req %s\n", addr)
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("myhttp res body:%s\n", body)
	fmt.Println(resp.StatusCode)
	if resp.StatusCode == 200 {
		fmt.Println("myhttp ok")
	} else {
		fmt.Printf("myhttp StatusCode:%d\n", resp.StatusCode)
	}
}
