package server

import (
	"fmt"
	"net/http"
)

func ServerHttpInit() error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world!\n")
	})
	return http.ListenAndServe("127.0.0.1:8082", nil)
}
