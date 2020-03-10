package myhttp

import (
	"fmt"
	"io"
	"lession/mysrclib"
	tools "lession/tools/mytime"
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
