package mysrclib

import (
	"fmt"
	"io"
	"io/ioutil"
	tools "lesson/tools/mytime"
	"log"
	"os"
)

var (
	MyTrace   *log.Logger
	MyInfo    *log.Logger
	MyWarning *log.Logger
	MyError   *log.Logger
	//MyFile			*log.Logger
)

const (
	PREFIX_TRACE   = "MyTrace:   "
	PREFIX_INFO    = "MyInfo:    "
	PREFIX_WARNING = "MyWarning: "
	PREFIX_ERROR   = "MyError:   "
	//PREFIX_FILE		= ""
	FLAG = log.Ldate | log.Ltime | log.Lshortfile
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func DirCreate(path string) (err error) {
	err = nil
	var exist = true

	exist, err = PathExists(path)
	if err != nil {
		return
	}
	if exist {
		return
	} else {
		err = os.Mkdir(path, 0777)
		if err != nil {
			return
		}
		exist = true
	}

	return
}

func init() {
	fmt.Println("---------welcome mylogger init----------")

	path := "./log/"

	if err := DirCreate(path); err != nil {
		fmt.Printf("test mkdir path failed,err:%v\n", err)
		return
	}

	fileName := path + "errors_" + tools.GetFileDateName_Day()
	fileName += ".log"
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("test open file failed,err:%v\n", err)
		return
	}

	MyTrace = log.New(ioutil.Discard, PREFIX_TRACE, FLAG)
	MyInfo = log.New(os.Stdout, PREFIX_INFO, FLAG)
	MyWarning = log.New(os.Stdout, PREFIX_WARNING, FLAG)
	MyError = log.New(io.MultiWriter(file, os.Stderr), PREFIX_ERROR, FLAG)

}

func TestMylogger() {
	fmt.Println("---------test muti interface----------")
	MyTrace.Println("I have something standard to say")
	MyInfo.Println("Special MyInformation")
	MyWarning.Println("There is something you need to know about")
	//MyError.Println("Something has failed")
}
