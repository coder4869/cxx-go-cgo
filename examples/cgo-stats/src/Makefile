# 配置项目目录
SRC_DIR  := $(shell pwd)
export LD_LIBRARY_PATH=$(SRC_DIR)/stats

# 初始化命令变量
GOROOT  := /usr/local/go
GO      := $(GOROOT)/bin/go
GOPATH  := $(shell $(GO) env GOPATH)
GOMOD   := $(GO) mod
GOBUILD := $(GO) build
GOTEST  := $(GO) test
GOPKGS  := $$($(GO) list ./...| grep -vE "vendor")

# 执行编译：使用命令 make 或 make all 执行, 执行顺序 prepare -> compile -> test -> package
all: prepare test
# prepare 准备阶段。单独执行命令: make prepare
prepare:
	git config --global http.sslVerify false # 设置 git，保证 github mirror 能够下载
	# 开启go mod模式
	$(GO) env -w GO111MODULE="on"                                   
	## 配置 GONOSUMDB，工程中部分代码库不支持 sumdb 索引，暂时屏蔽此功能
	$(GO) env -w GONOSUMDB=\*
	## 配置 GOPROXY 为七牛云代理，可以下载墙外代码
	$(GO) env -w GOPROXY=https://goproxy.cn,direct         
	$(GO) env
	
# test 测试阶段。单独执行命令: make test
test: 
	$(GO) clean -modcache
	# $(GOMOD) tidy
	$(GOMOD) download # 下载Go依赖
	# 输出测试结果到文件中
	$(GOTEST) -v -cover $(GOPKGS) > stats_test_log.out
