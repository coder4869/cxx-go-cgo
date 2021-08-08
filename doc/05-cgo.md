## cgo
Go 语言通过自带的一个叫 CGO 的工具来支持 C 语言函数调用，同时我们可以用 Go 语言导出 C 动态库接口给其它语言使用。

### go 调用 C
go 调用 C 主要包括三部分：引用的 C 资源声明；启用 CGO 能力；具体方法调用。

#### 引用的 C 资源声明
在调用 C 库的 go 源码文件中，以注释的形式声明引用的 C 资源信息。引用资源配置可以分为两种类型：

- 编译、链接等选项指定。以 `#cgo` 开头，配置方式与 Make 指令的 C 编译参数配置基本一致。比如：
    * 使用 `CFLAGS` 配置编译选项。如：常用的 `-I` 指定头文件依赖路径。
    * 使用 `LDFLAGS` 配置链接选项。如：使用 `-Wl,-rpath-link` 指定依赖库路径。若有多个依赖库目录，可以写多行配置。
    * 其中 `${SRCDIR}` 是系统自带变量，为当前 go 源代码文件目录的绝对路径。在使用时，尽量保持最终的 bin 可执行文件与依赖库的相对位置，与调用C代码的go文件一致。否则需要修正最终的可执行 bin 文件与 依赖库的位置关系。
- 具体的头文件内件引入。与 C/C++ 语言相同，直接使用 `#include` 指令即可。

示例如下 (正式代码中 `--` 及其后面的内容需要删除)：

```go
//////////////////////////// 目录架构一 ////////////////////////////
//// 源码目录层级（cgo.go 与 include、lib、thirdparty 在同级目录下）
// ${SRCDIR} ---  cgo.go (调用C代码的 go 文件)
//           |--  include (依赖的头文件目录)
//           |--  lib (go 直接依赖库目录)
//           |--  thirdparty/lib (go 间接依赖库目录)
//// 产物目录层级 (相对目录架构二，bin 无需修正，可直接执行)
// ${SRCDIR} ---  bin (最终的可执行文件)
//           |--  include (依赖的头文件目录)
//           |--  lib (go 直接依赖库目录)
//           |--  thirdparty/lib (go 间接依赖库目录)
//// cgo.go 文件中编译、链接等选项配置
/*
#cgo CFLAGS: -I./include                                    -- 编译选项，声明头文件所在文件夹路径
#cgo LDFLAGS: -Wl,-rpath-link="${SRCDIR}/lib"               -- 链接选项，go 直接依赖库目录，-rpath-link 参数，在 mac 上不识别，在 linux 上可识别。
#cgo LDFLAGS: -Wl,-rpath-link="${SRCDIR}/thirdparty/lib"    -- 链接选项，go 间接依赖库目录
#cgo LDFLAGS: -L${SRCDIR}/lib -lc_interface     -- 链接 ${SRCDIR}/lib 文件夹下的 libc_interface.so 库作为直接调用的 C-API 入口

#include <stdlib.h>                                     -- 引入 C 标准头文件
#include "c_interface.h"                                -- 引入自定义的 C 头文件
*/
import "C"                                              -- 启用 CGO 能力，与上面的 C 声明之间不能包含空格。


//////////////////////////// 目录架构二 ////////////////////////////
//// 源码目录层级（${SRCDIR} 文件夹与 deps 在同级目录下，cgo.go 与 include、lib、thirdparty 不在同级目录下）
// ${PROJECT_DIR} --- ${SRCDIR} ---  cgo.go (调用C代码的 go 文件)
//                 |  deps      ---
//                              |--  include (依赖的头文件目录)
//                              |--  lib (go 直接依赖库目录)
//                              |--  thirdparty/lib (go 间接依赖库目录)
//// 产物目录层级 (相对目录架构一，bin 需修正与 lib 中依赖库的依赖关系，才可执行，否则会报错连接库找不到)
// ${SRCDIR} ---  bin (最终的可执行文件)
//           |--  include (依赖的头文件目录)
//           |--  lib (go 直接依赖库目录)
//           |--  thirdparty/lib (go 间接依赖库目录)
//// cgo.go 文件中编译、链接等选项配置
/*
#cgo CFLAGS: -I../deps/include                                   -- 编译选项，声明头文件所在文件夹路径
#cgo LDFLAGS: -Wl,-rpath-link="${SRCDIR}/../deps/lib"            -- 链接选项，go 直接依赖库目录
#cgo LDFLAGS: -Wl,-rpath-link="${SRCDIR}/../deps/thirdparty/lib" -- 链接选项，go 间接依赖库目录
#cgo LDFLAGS: -L${SRCDIR}/../deps/lib -lc_interface     -- 链接 ${SRCDIR}/lib 文件夹下的 libc_interface.so 库作为直接调用的 C-API 入口

#include <stdlib.h>                                     -- 引入 C 标准头文件
#include "c_interface.h"                                -- 引入自定义的 C 头文件
*/
import "C"                                              -- 启用 CGO 能力，与上面的 C 声明之间不能包含空格。
```

#### 启用 CGO 能力
- 开启 CGO 能力：在调用 C 库的 go 源码文件中添加 `import "C"` 语句 (参见上面示例)。且 import 和上边的 C 声明之间不能有空格。

#### 具体方法调用
参见 examples/cgo-base/src/cgo/cgo-user-def.go

### 参考资料
- CGO编程：https://www.cntofu.com/book/73/ch2-cgo/readme.md
