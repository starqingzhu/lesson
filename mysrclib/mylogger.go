package mysrclib

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	MyTrace 		*log.Logger
	MyInfo		*log.Logger
	MyWarning		*log.Logger
	MyError		*log.Logger
)

const (
	PREFIX_TRACE	= "MyTrace:   "
	PREFIX_INFO		= "MyInfo:    "
	PREFIX_WARNING	= "MyWarning: "
	PREFIX_ERROR	= "MyError:   "
	FLAG	= log.Ldate | log.Ltime | log.Lshortfile
)

func  init(){
	fmt.Println("---------welcome mylogger init----------")
	file ,err := os.OpenFile("./errors.txt",os.O_CREATE | os.O_WRONLY | os.O_APPEND,0666)
	if err != nil{
		fmt.Printf("open file failed\n")
		return
	}
	fmt.Printf("open file success\n")

	MyTrace	= log.New(ioutil.Discard,PREFIX_TRACE,FLAG)
	MyInfo	= log.New(os.Stdout,PREFIX_INFO,FLAG)
	MyWarning	= log.New(os.Stdout,PREFIX_WARNING,FLAG)
	MyError	= log.New(io.MultiWriter(file,os.Stderr),PREFIX_ERROR,FLAG)
}

func TestMylogger(){
	fmt.Println("---------test muti interface----------")
	MyTrace.Println("I have something standard to say")
	MyInfo.Println("Special MyInformation")
	MyWarning.Println("There is something you need to know about")
	MyError.Println("Something has failed")
}