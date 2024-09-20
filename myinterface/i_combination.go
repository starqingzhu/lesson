package myinterface

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

//结构体
type (
	LimitedReader struct {
		R Reader
		N int64
	}
)

func (l *LimitedReader) Read(p []byte) (n int, err error) {
	return 0, nil
}

func LimitReader(r Reader, n int64) Reader {
	return &LimitedReader{r, n}
}

/*
wraper
*/
type (
	capitalizedReader struct {
		r io.Reader
	}
)

func (r *capitalizedReader) Read(p []byte) (int, error) {
	n, err := r.r.Read(p)
	if err != nil {
		return 0, err
	}

	q := bytes.ToUpper(p)
	for i, v := range q {
		p[i] = v
	}

	return n, nil
}

func CapReader(r io.Reader) io.Reader {
	return &capitalizedReader{
		r: r,
	}
}

func PrintInterfaceInterCombination() {
	r := strings.NewReader("hello, gopher!\n")
	r1 := CapReader(r)

	if _, err := io.Copy(os.Stdout, r1); err != nil {
		log.Fatal(err)
	}
}

//适配器
type (
	Handler interface {
		ServeHTTP(http.ResponseWriter, *http.Request)
	}

	//函数签名
	HandlerFunc func(http.ResponseWriter, *http.Request)
)

//func (f HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	f(w, r)
//}
func greetings(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!")
}

func PrintInterfaceHandler() {
	http.ListenAndServe(":8090", http.HandlerFunc(greetings))
}
