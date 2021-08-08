## go-start
- 文档资料：
    * 英文版：https://golang.org/doc/
    * 中文版：https://go-zh.org/doc/
    * wiki：https://github.com/golang/go/wiki
    * Go语言高级编程：https://www.cntofu.com/book/73/index.html

### 基础资料
- 1. 安装软件包：https://golang.org/doc/install
- 2. 熟悉相关基础概念：(GOPATH、GOROOT等)环境变量、工作空间、编译配置、包管理、测试等。https://golang.org/doc/install
- 3. 资源包与功能：https://go-zh.org/pkg/
- 4. 基本语法：https://go-zh.org/doc/effective_go.html
- 5. 编程规范：https://go-zh.org/ref/spec
- 6. 常见问题解答：https://go-zh.org/doc/faq
- 7. 内存模型：https://go-zh.org/ref/mem
- 8. 命令文档：https://go-zh.org/doc/cmd
- 9. 发行历史：https://go-zh.org/doc/devel/release.html
- 10. wiki-Learn：https://github.com/golang/go/wiki/Learn

### 扩展资料 - 文章
- 1. 官方博客：https://blog.golang.org
- 2. (codewalk)代码示例：
    * 函数 - Go 语言中的一等公民：https://go-zh.org/doc/codewalk/functions/
    * 生成任意文本：马尔可夫链算法：https://go-zh.org/doc/codewalk/markov/
    * 通过通信共享内存：https://go-zh.org/doc/codewalk/sharemem/
    * 编写Web应用 - 构建简单的Web应用：https://go-zh.org/doc/articles/wiki/
- 3. 语言
    * JSON-RPC：有关接口的故事 - https://go-zh.org/blog/json-rpc-tale-of-interfaces
    * Go 的声明语法 - https://go-zh.org/blog/gos-declaration-syntax
    * Defer、Panic 和 Recover - https://go-zh.org/blog/defer-panic-and-recover
    * Go 并发模式：超时，继续 - https://go-zh.org/blog/go-concurrency-patterns-timing-out-and
    * Go 切片：用法和本质 - https://go-zh.org/blog/go-slices-usage-and-internals
    * GIF 解码器：Go 接口练习 - https://go-zh.org/blog/gif-decoder-exercise-in-go-interfaces
    * Go 与错误处理 - https://go-zh.org/blog/error-handling-and-go
    * 组织 Go 代码 - https://go-zh.org/blog/organizing-go-code
- 4. 包
    * [JSON 和 Go](https://go-zh.org/blog/json-and-go) - 使用[json](https://go-zh.org/pkg/encoding/json/)包。
    * [数据一坨](https://go-zh.org/blog/gobs-of-data) - [gob](https://go-zh.org/pkg/encoding/gob/)包的设计与使用。
    * [反射法则](https://go-zh.org/blog/laws-of-reflection) - [reflect](https://go-zh.org/pkg/reflect/)包基础。
    * [Go 图像包](https://go-zh.org/blog/go-image-package) - [image](https://go-zh.org/pkg/image/)包基础。
    * [Go 图像绘制包](https://go-zh.org/blog/go-imagedraw-package) - [image/draw](https://go-zh.org/pkg/image/draw/)包基础。
- 5. 工具
    * [关于 Go 命令](https://go-zh.org/doc/articles/go_command.html) - 为什么写它？它是什么？它不是什么？它怎么用？
    * [C? Go? Cgo!](https://go-zh.org/blog/c-go-cgo) - 使用[cgo](https://go-zh.org/cmd/cgo/)连接 C 代码。
    * [使用GDB调试Go代码](https://go-zh.org/doc/gdb)
    * [Godoc：编写 Go 代码文档](https://go-zh.org/blog/godoc-documenting-go-code) - 为[godoc](https://pkg.go.dev/github.com/Go-zh/tools/cmd/godoc)编写好的文档。
    * [Go 程序性能分析](https://go-zh.org/blog/profiling-go-programs)
    * [数据竞态检测器](https://go-zh.org/doc/articles/race_detector.html) - 测试竞态条件下的 Go 程序。
    * [Go 竞态检测器介绍](https://go-zh.org/blog/race-detector) - 对竞态检测器的介绍。
    * [Go 汇编器快速指南](https://go-zh.org/doc/asm) - 对 Go 使用的汇编的介绍。
- 6. 更多资料
    * [wiki-Articles](https://github.com/golang/go/wiki/Articles) 

### 扩展资料 - 演讲
- [Go 视频教程](http://research.swtch.com/gotour)
    * 有三样东西使得 Go 快速、有趣而高效：接口、反射与并发。我们将通过构建一个玩具网络爬虫来展示它们。

- [代码的优雅增长](https://vimeo.com/53221560)
    * Go 的关键设计目标之一就是代码的适应性，它应当易于得到简单的设计，并以干净而自然的方式构建。 在此演讲中，Andrew Gerrand 描述了一种简单的“轮流聊天”服务，它匹配一对 TCP 接入的连接， 并使用 Go 的并发机制、接口和标准库来为它扩展出 Web 界面和其它特性。当该程序的函数戏剧性地改变后， Go 的灵活性在它增长时保留了其原始的设计。

- [Go 的并发模式](https://www.youtube.com/watch?v=f6kdp27TYZs)
    * 并发是设计高性能网络服务的关键。Go 的并发原语（Go程与信道）提供了一个表达并发实行的简单而高效的手段。在此演讲中，我们将看见如何用简单的 Go 代码优雅地解决棘手的并发问题。

- [Go 的高级并发模式](https://www.youtube.com/watch?v=QDDwwePbDtw)
    * 此演讲是对 Go 的并发模式的进一步扩充，它对 Go 的并发原语进行了更加深入的探讨。

- 更多演讲
    * Go 演讲网站：http://talks.zh-golang.appspot.com
    * wiki-GoTalks：https://github.com/golang/go/wiki/GoTalks

### 扩展资料 - 非英语文档
    * https://github.com/golang/go/wiki/NonEnglish
