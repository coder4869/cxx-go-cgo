# 配置项目目录
OUTPUT   := build
SRC_DIR  := $(shell pwd)
HOME_DIR := $(SRC_DIR)/..
OUT_DIR  := $(HOME_DIR)/$(OUTPUT)
LIB_ROOT := $(HOME_DIR)/library

# 可执行文件名称，可自行修改
APP_NAME := cgo_base
# PATH 环境变量添加 CentOS 6.3 自带的 gcc 版本号
PATH_OLD := $(PATH)
export PATH=/opt/compiler/gcc-4.8.2/bin:$(PATH_OLD)
# export CFLAGS=$(LIB_ROOT)/include
# export C_INCLUDE_PATH=$(LIB_ROOT)/include
# export LD_LIBRARY_PATH=$(LIB_ROOT)/lib

# 初始化命令变量
GOROOT  := /usr/local/go
GO      := $(GOROOT)/bin/go
GOPATH  := $(shell $(GO) env GOPATH)
GOMOD   := $(GO) mod
GOBUILD := $(GO) build
GOTEST  := $(GO) test
GOPKGS  := $$($(GO) list ./...| grep -vE "vendor")

# 执行编译：使用命令 make 或 make all 执行, 执行顺序 prepare -> compile -> test -> package
all: prepare compile package
# prepare 准备阶段。单独执行命令: make prepare
prepare: prepare-dep
prepare-dep:
	git config --global http.sslVerify false # 设置 git，保证 github mirror 能够下载
set-env:
	# 开启go mod模式
	$(GO) env -w GO111MODULE="on"                          
	# 配置 GONOPROXY 环境变量，所有百度内代码，不走代理
	$(GO) env -w GONOPROXY=\*\*.baidu.com\*\*              
	## 配置 GONOSUMDB，工程中部分代码库不支持 sumdb 索引，暂时屏蔽此功能
	$(GO) env -w GONOSUMDB=\*
	## 配置 GOPROXY 为七牛云代理，可以下载墙外代码
	$(GO) env -w GOPROXY=https://goproxy.cn,direct         
	$(GO) env

# complile 编译阶段。单独执行命令: make compile
compile: build
build: set-env
	$(GO) clean -modcache
	# $(GOMOD) tidy
	$(GOMOD) download #下载Go依赖
	mkdir -p $(OUT_DIR)
	cd $(HOME_DIR) && bash scripts/lib_process.sh copy
	$(GOBUILD) -o $(APP_NAME)
# test 测试阶段。单独执行命令: make test
test: test-case
test-case: set-env
	$(GOTEST) -v -cover $(GOPKGS)
# 与覆盖率平台打通，输出测试结果到文件中
# @$(GOTEST) -v -json -coverprofile=coverage.out $(GOPKGS) > testlog.out
# package 打包阶段：打包编译产出到 $(OUT_DIR) 目录。单独执行命令: make package
package: package-bin
package-bin:
	$(shell rm -rf $(OUT_DIR))
	$(shell mkdir -p $(OUT_DIR)/bin)
	$(shell mkdir -p $(OUT_DIR)/bin/lib)

	$(shell cp    $(APP_NAME)		$(OUT_DIR)/bin/)
	$(shell cp -r $(LIB_ROOT)/lib/	$(OUT_DIR)/bin/lib/)
	$(shell cp -r conf/*			$(OUT_DIR)/bin/conf/)
	cd $(OUT_DIR) && tar cvzf $(APP_NAME).tar.gz bin 
	# rm -rf $(OUT_DIR)/bin

# clean 清理阶段，清除过程中的输出。单独执行命令: make clean
clean:
	rm -rf $(OUT_DIR) $(APP_NAME)
# avoid filename conflict and speed up build
.PHONY: all prepare compile test package clean build