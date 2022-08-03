package mysync

import (
	"bytes"
	"sync"
	"testing"
)

var bufPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func writeBufFromPool(data string) {
	b := bufPool.Get().(*bytes.Buffer)
	b.Reset()
	b.WriteString(data)
	bufPool.Put(b)
}

func writeBuffFromNew(data string) *bytes.Buffer {
	b := new(bytes.Buffer)
	b.WriteString(data)
	return b
}

func BenchmarkWithoutPool(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		writeBuffFromNew("hello")
	}
}

func BenchmarkWithPool(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		writeBufFromPool("hello")
	}
}
