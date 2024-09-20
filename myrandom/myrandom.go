package myrandom

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	recover()
	fmt.Println("---------welcome myrandom init----------")
}

func PrintRandom() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 3; i++ {
		fmt.Printf("myrandom %d\n", rand.Intn(100))
	}
}
