package mysrclib

import (
	"fmt"
	"log"
)

const (
	Prefix	= ""
	Flags	= log.Ldate | log.Ltime | log.Lshortfile
)

func init(){
	log.SetPrefix(Prefix)
	log.SetFlags(Flags)
}

func TestLogFmt(){
	fmt.Printf("-----------test logfmt ------------\n")
	log.Printf("hello world\n")
}
