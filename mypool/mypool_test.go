package mypool

import (
	"bytes"
	"encoding/json"
	"testing"
)

func BenchmarkUnmarsh(b *testing.B) {
	temp := &Student{}
	var buf, _ = json.Marshal(temp)

	for n := 0; n < b.N; n++ {
		stu := &Student{}
		json.Unmarshal(buf, stu)
	}
}

func BenchmarkGetObjPool(b *testing.B) {
	studentPool := GetObjPool()
	temp := studentPool.Get()
	var buf, _ = json.Marshal(temp)

	for n := 0; n < b.N; n++ {
		stu := studentPool.Get()
		json.Unmarshal(buf, stu)
		studentPool.Put(stu)
	}
}

func BenchmarkBufferWithPool(b *testing.B) {
	bufferPool := GetBuffPool()
	data := make([]byte, 10000)

	for n := 0; n < b.N; n++ {
		buf := bufferPool.Get().(*bytes.Buffer)
		buf.Write(data)
		buf.Reset()
		bufferPool.Put(buf)
	}
}

func BenchmarkBuffer(b *testing.B) {
	data := make([]byte, 10000)
	for n := 0; n < b.N; n++ {
		var buf bytes.Buffer
		buf.Write(data)
	}
}
