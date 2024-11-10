package main

import (
	"fmt"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}

func main() {
    http.HandleFunc("/", helloWorld) // 设置访问的路由以及其处理函数
    http.ListenAndServe(":8080", nil) // 设置监听的端口，并启动服务
}