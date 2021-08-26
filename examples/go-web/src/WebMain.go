// Copyright (c) 2016 coder4869 ( https://github.com/coder4869/cxx-go-cgo/tree/main/examples). All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
	"net/http"

	"goweb/routers"
)

func main() {
	// log setting
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("start WebServe8080 v2.0 beta")

	routers.Init()

	// listen port setting
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
