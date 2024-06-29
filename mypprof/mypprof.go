package mypprof

import (
	"log"
	"os"
	"path/filepath"
	"runtime/pprof"
	"time"
)

func CreateProfileFile() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalln("get current directory failed.", err)
	}

	fileName := filepath.Join(dir, "mypprof", "profile", "heap")
	f, err1 := os.Create(fileName)
	if err1 != nil {
		log.Fatalln("Create current directory failed.", err1)
	}
	_ = pprof.StartCPUProfile(f)

	defer pprof.StopCPUProfile()
	for i := 0; i < 50; i++ {
		log.Println("sleep ", i)
		time.Sleep(100 * time.Millisecond)
	}
}
