// Copyright (c) 2016 coder4869 ( https://github.com/coder4869/cxx-go-cgo/tree/main/examples). All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package controllers

import (
	"net/http"

	"github.com/coder4869/golibs/glio"
)

func Index(w http.ResponseWriter, r *http.Request) {
	glio.FFLPrintf(w, "Index Page is in developing\n") //output to client with w
}
