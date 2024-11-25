package main

import (
	"fmt"
	"hello-go/etcd"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	etcd.Init()
	go etcd.WatchKey("test")

	for i := 0; i < 5; i++ {
		go etcd.TryLock(i, "test/1")
	}

	http.HandleFunc("/", helloWorld)  // 设置访问的路由以及其处理函数
	http.ListenAndServe(":8080", nil) // 设置监听的端口，并启动服务
}
