package cgo

/*
#cgo CFLAGS: -I../../library/include
#cgo windows LDFLAGS: 
#cgo darwin LDFLAGS: -Wl,-rpath,${SRCDIR}/../../library/lib
#cgo linux LDFLAGS: -Wl,-rpath-link="${SRCDIR}/../../library/lib" // -rpath-link 参数，在 mac 上不识别，在 linux 上可识别。
#cgo linux LDFLAGS: -Wl,-rpath,'$ORIGIN'
#cgo LDFLAGS: -L${SRCDIR}/../../library/lib -lCBase

#include <stdlib.h>
#include "cBase/c_base/base_type.h"
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