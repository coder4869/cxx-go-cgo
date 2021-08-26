// Copyright (c) 2016 coder4869 ( https://github.com/coder4869/GoWeb ). All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package web

import (
	"testing"
)

func TestUserLoginServe(t *testing.T) {
	HttpGet()
	HttpPost()
}

// Parallel Test
func Benchmark_UserLoginServeParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			go HttpGet()
		}
	})
}

func Benchmark_UserLoginServe(b *testing.B) {
	// must run b.N times. b.N will auto-adjust in running,
	// this ensure both time cost and caculated test data is reasonable.
	for i := 0; i < b.N; i++ {
		go HttpPost()
	}
}
