// Copyright (c) 2016 coder4869 ( https://github.com/coder4869/GoWeb ). All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package web

import (
	"bytes"
	"fmt"
	"log"
	"os"
    "io/ioutil"
    "net/http"
)

func HttpGet()  {
	urlString := "http://127.0.0.1:8080/goweb/user/login?username=zhangsan&password=123456"

	// Get请求
	res, err := http.Get(urlString)
	if err != nil {
		log.Fatal(err)
	}
	// 利用 ioutil 包读取服务器返回的数据
	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()//一定要记得关闭连接
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s", data)
}

func HttpPost()  {
	urlString := "http://127.0.0.1:8080/goweb/user/login"

	postData := "{" +
					"\"username\"	: \"zhangsan\", " +
					"\"password\"	: \"123456\"" +
				"}"

	var jsonStr = []byte(postData)
	// fmt.Println("jsonStr", bytes.NewBuffer(jsonStr))

	req, err := http.NewRequest("POST", urlString, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil{
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("status", resp.Status)
	fmt.Println("response:", resp.Header)
	// 利用 ioutil 包读取服务器返回的数据
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

func saveStringToFile(path string, data string)  {
	fi, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0644)

    if err != nil {
        panic(err)
    }
    defer fi.Close()

	fi.WriteString(data)
}