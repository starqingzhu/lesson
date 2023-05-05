package mylist

import (
	"container/list"
	"container/ring"
	"fmt"
)

func init() {
	fmt.Println("---------welcome mylist init----------")
}

func TestMyList() {
	var l list.List
	l.Init()

	var r ring.Ring
	_ = r.Len()
}
