package controller

import (
	"net/http"
	"chineseToPinAndWubi/sys/service"
	"encoding/json"
)
type ResponseJson struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
var response ResponseJson
func GetWbPyHandler (w http.ResponseWriter,r *http.Request){
	//设置头信息，防止跨域问题
	SetHeader(w)
	r.ParseForm()
	defer r.Body.Close()
	v,ok := r.Form["strs"]
	if !ok {

	}
	if r.Method == "GET"{
		r.ParseForm()
		data:=service.GetWbPyService(v[0])
		response.Code = "100000"
		response.Msg = "不知此的url, 最可能的原因是url拼写错误!"
		responseJsnByte, _ := json.Marshal(response)
		w.Write(responseJsnByte)
		return
	}else {
		response.Code = "100000"
		response.Msg = "不知此的url, 最可能的原因是url拼写错误!"
		responseJsnByte, _ := json.Marshal(response)
		w.Write(responseJsnByte)
		return
	}
}
