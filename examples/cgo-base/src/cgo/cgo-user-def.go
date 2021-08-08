package cgo

/*
#cgo CFLAGS: -I../../library/include
// #cgo LDFLAGS: -Wl,-rpath-link="${SRCDIR}/../../library/lib" // -rpath-link 参数，在 mac 上不识别，在 linux 上可识别。
#cgo LDFLAGS: -L${SRCDIR}/../../library/lib -lCXXBase

#include <stdlib.h>
#include "base/base_type.h"
*/
import "C"

import (
	"fmt"
	// "unsafe"
)

func CgoUserDef()  {
	ret_add := C.base_add(2, 3);
	fmt.Println("2 + 3 =", ret_add);
}