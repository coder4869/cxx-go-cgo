#!/bin/bash
# protoc 下载地址：https://github.com/protocolbuffers/protobuf/releases

script_dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd ${script_dir}

protoc_dir=${script_dir}/protoc-3.17.3/bin-x86_64

cache_dir=${script_dir}/cache
pb_tool_dir=${cache_dir}/protobuf/protoc-gen-go
function prepare_protobuf() {
    mkdir -p ${cache_dir} 
    echo ${pb_tool_dir}

    if [ ! -d "${pb_tool_dir}" ]; then
        # unzip local first
        if [ -f "protobuf.tar.gz" ]; then
            tar -zvxf ${script_dir}/protobuf.tar.gz -C ${cache_dir}
        fi 

        # unzip failed -- get from github
        if [ ! -d "${pb_tool_dir}" ]; then
            cd ${cache_dir}
            git clone https://github.com/golang/protobuf.git
            cd -
        fi
    fi
}

function gen_protoc_tools() {
    pb_tool_file=${pb_tool_dir}/protoc-gen-go

    if [ -d "${pb_tool_dir}" ]; then
        cd ${pb_tool_dir}
        # if [ ! -f "protoc-gen-go" ]; then
            go build -o protoc-gen-go
        # fi
        cd -
    fi

    mkdir -p ${protoc_dir} && cp ${pb_tool_file} ${protoc_dir}/
}

# generate "pb/*.pb.go"
src_dir=${script_dir}/../../examples/cgo-pb/src
function gen_pb_go() {
    # "$(uname)" == "Darwin"
    protoc_file=${protoc_dir}/protoc-osx
    if [ "$(uname)" == "Linux" ]; then
        protoc_file=${protoc_dir}/protoc-linux
    fi

    export PATH=${protoc_dir}:$PATH
    ${protoc_file} --proto_path=${src_dir} --go_out=${src_dir} pb/req/request.proto
    ${protoc_file} --proto_path=${src_dir} --go_out=${src_dir} pb/resp/response.proto
}

function clean_cache() {
    rm -rf ${cache_dir}
}

prepare_protobuf
gen_protoc_tools
gen_pb_go
clean_cache


