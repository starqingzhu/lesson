package mysrclib

import (
	"bytes"
	"fmt"
	"os"
)

func TestMyIo(){
	fmt.Printf("-----------test myIo ------------\n")

	var b bytes.Buffer
	b.WriteString("hello ")
	fmt.Fprintf(&b,"%s%s","world!"," !!!!\n")

	b.WriteTo(os.Stdout)
}