package main

import (
	"net/http"
	"chineseToPinAndWubi/sys/controller"
	"flag"
	"encoding/json"
	"fmt"
)
type ResponseJson struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
var response ResponseJson
/*
	main()
	程序运行的入口
*/

func main() {
	//连接数据库

	flag.Parse()

	//路由转发
	http.HandleFunc("/", DefaultHandler)
	http.HandleFunc("/his/sys/get/wbpy", controller.GetWbPyHandler)//获取数据库时间


	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("端口监听错误!")
		return
	}
}

func DefaultHandler(w http.ResponseWriter, r *http.Request) {

	response.Code = "100000"
	response.Msg = "不知此的url, 最可能的原因是url拼写错误!"
	responseJsnByte, _ := json.Marshal(response)
	w.Write(responseJsnByte)
	return
}