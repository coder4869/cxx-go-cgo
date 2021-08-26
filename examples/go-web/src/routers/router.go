// Copyright (c) 2016 coder4869 ( https://github.com/coder4869/cxx-go-cgo/tree/main/examples). All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package routers

import (
	"net/http"

	"goweb/controllers"
)

func Init()  {
	// router setting
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/goweb/user/login", controllers.UserLoginServe)
}