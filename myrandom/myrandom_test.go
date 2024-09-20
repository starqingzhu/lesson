package myrandom

import (
	"fmt"
	"testing"
)

func TestRandom(t *testing.T) {
	t.Parallel()
	PrintRandom()
}

func BenchmarkPrintRandom(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			fmt.Println("xxxx")
		}
	})
}
