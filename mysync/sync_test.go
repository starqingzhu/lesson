package mysync

import (
	"sync"
	"testing"
)

var cs = 0
var mu sync.Mutex
var c = make(chan struct{}, 1)

func criticalSectionSyncByMutex() {
	mu.Lock()
	defer mu.Unlock()
	cs++
}
func criticalSectionSyncByChan() {
	c <- struct{}{}
	defer func() {
		<-c
	}()
	cs++
}

func BenchmarkCriticalSectionSyncByMutex(b *testing.B) {
	for n := 0; n < b.N; n++ {
		criticalSectionSyncByMutex()
	}
}

func BenchmarkCriticalSectionSyncByChan(b *testing.B) {
	for n := 0; n < b.N; n++ {
		criticalSectionSyncByChan()
	}
}
