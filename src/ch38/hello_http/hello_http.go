package main

import (
	"fmt"
	"net/http"

	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})
	http.HandleFunc("/time/", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		timeStr := fmt.Sprintf("{\"time\": \"%s\"}", t)
		// go 里隐式类型转换会报错，只能使用强制类型转换
		w.Write([]byte(timeStr))
	})

	http.ListenAndServe(":8080", nil)
}