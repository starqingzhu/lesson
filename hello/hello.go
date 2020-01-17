package hello

import (
	"fmt"
	//"github.com/spf13/viper"
)

func init(){
	fmt.Println("---------welcome hello init----------")
}

func PrintHello() {
	fmt.Println("package hello world")
}
