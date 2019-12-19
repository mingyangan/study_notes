package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)

var datas []string

func add(str string) string {
	data := []byte(str)
	sData := string(data)
	datas = append(datas, sData)

	return sData
}

func main() {
	go func() {
		for {
			add("https://github.com/")
			time.Sleep(time.Millisecond)
		}
	}()

	fmt.Println("Server Started")
	http.ListenAndServe("0.0.0.0:6060", nil)
}
