package main

import (
	"encoding/json"
	"fmt"
	"hello-go/configs"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	data, _ := json.Marshal(configs.Get())
	fmt.Fprintf(w, "111:"+string(data))
}

func main() {
	http.HandleFunc("/", helloWorld)  // 设置访问的路由以及其处理函数
	http.ListenAndServe(":8080", nil) // 设置监听的端口，并启动服务
}
