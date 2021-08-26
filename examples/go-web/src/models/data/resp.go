// Copyright (c) 2016 coder4869 ( https://github.com/coder4869/cxx-go-cgo/tree/main/examples). All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package data

var (
	AesKey      string = "1234567812345678"
)

// Request Return Info
type RespInfo struct {
	ErrCode	int32
	ErrMsg 	string
	Data	string
}
