# cxx-go-cgo
go 相关的学习资料，以及基于 cgo  的  c++ 与 go 交互 examples 等。


## doc
存放调研文档资料, 做好分类整理工作。


## examples
测试 demos, 做好 demo 的按文件夹分类整理工作。文件夹名称要见名知义。

### examples/cgo-cxx
c++ cmake project for go.

### examples/cgo-base
cgo project : with deps of project cgo-cxx.

https://insujang.github.io/2019-12-10/cmake-for-custom-library-build-system-in-go/

### examples/cgo-pb
cgo project with protocbuf. 

项目中的 *.pb.go 文件，是 protoc 工具调用 protoc-gen-go 插件，根据 *.proto 文件自动生成的，无需修改。生成过程参见，tools下的protoc.sh

- protoc 下载地址：https://github.com/protocolbuffers/protobuf/releases
- protoc-gen-go 插件源码：https://github.com/golang/protobuf 【下载后，在 protoc-gen-go 目录执行编译】

### examples/cgo-stats
- 参考资料：https://github.com/akhenakh/statgo


### examples/go-web
- 第三方库：https://www.zhihu.com/question/27370112

## 参考资料
### GO官方资料
- 官网：https://golang.org/ 包含文档资料，安装包下载，在线编译环境等。
- 安装：通过 https://golang.org/dl/ 下载安装包，安装后通过 `$ go version` 查看版本
- 卸载：执行 `$ go env` 指令，查看安装目录, 删除 `GOROOT` 变量指向目录。
    * 比如：`GOROOT="/usr/local/go"`。则用 `$ sudo rm -rf /usr/local/go` 删除安装目录
- 学习文档：
    * 英文版：https://golang.org/doc/
    * 中文版-1：https://github.com/Go-zh/tour
    * 中文版-2：https://go-zh.org/doc/
    * wiki：https://github.com/golang/go/wiki


https://blog.icorer.com/index.php/archives/382/