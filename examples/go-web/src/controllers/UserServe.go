// Copyright (c) 2016 coder4869 ( https://github.com/coder4869/cxx-go-cgo/tree/main/examples). All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package controllers

import (
	"fmt"
	"encoding/json"
	"html"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/coder4869/golibs/glio"

	"goweb/libs"
	"goweb/models/data"
)

func UserLoginServe(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // 解析参数，默认是不会解析的
	
	if r.Method == "GET" {
		fmt.Println("method:", r.Method) // 获取请求的方法
		fmt.Println("username", r.Form["username"])
		fmt.Println("password", r.Form["password"])
		for k, v := range r.Form {
			fmt.Print("key:", k, "; ")
			fmt.Println("val:", strings.Join(v, ""))
		}
	} else if r.Method == "POST" {
		result, _:= ioutil.ReadAll(r.Body)
		r.Body.Close()
		fmt.Printf("%s\n", result)

		// 未知类型的推荐处理方法
		var f interface{}
		json.Unmarshal(result, &f)
		m := f.(map[string]interface{})
		for k, v := range m {
			switch vv := v.(type) {
				case string:
					fmt.Println(k, "is string", vv)
				case int:
					fmt.Println(k, "is int", vv)
				case float64:
					fmt.Println(k,"is float64",vv)
				case []interface{}:
					fmt.Println(k, "is an array:")
					for i, u := range vv {
						fmt.Println(i, u)
					}
				default:
					fmt.Println(k, "is of a type I don't know how to handle")
			}
		}
	}

	var resp = new(data.RespInfo) 

	resp.ErrCode = libs.RESULT_SUCCESS
	resp.ErrMsg = libs.ErrInfos[resp.ErrCode].ErrMsg
	resp.Data = "Current URL Path : " + html.EscapeString(r.URL.Path[1:])

	bytes, _ := json.Marshal(resp)
	glio.FFLPrintf(w, string(bytes) + "\n") // output to client with w
}
