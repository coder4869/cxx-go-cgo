#!/bin/bash

script_dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd ${script_dir}

function prepare_cmake_toolchains() {
    rm -rf ${script_dir}/cmake-toolchains
    # git clone https://github.com/coder4869/cmake-toolchains.git
    mkdir -p ${script_dir}/cmake-toolchains
    cp -rf ${script_dir}/../../../cmake-toolchains/* ${script_dir}/cmake-toolchains/
}

# generate "pb/*.pb.go"
function update_protobuf() {
    sh ${script_dir}/../../tools/protoc.sh
}

function update_cxx_library() {
    sh ${script_dir}/../cgo-cxx/build.sh
    cp -r ${script_dir}/../cgo-cxx/output/ ${script_dir}/library
}

################# build_go #################
src_dir=${script_dir}/src
bin_name=cgo_pb

function init_go_mod() {
    cd ${src_dir}
    if [ ! -f "go.mod" ]; then
        go mod init ${bin_name}
    fi

    if [ -f "go.mod" ]; then
       go mod tidy
    fi
}

build_dir=${script_dir}/build
function build_go() {
    mkdir -p ${build_dir} && rm -rf ${build_dir}/* 

    cd ${src_dir} && make
}

test_dir=${script_dir}/test/
function update_and_run_test() {
    mkdir -p ${test_dir} && rm -rf ${test_dir}/* 

    mv ${build_dir}/bin/* ${test_dir}
    sh ${script_dir}/scripts/lib_process.sh relink ${test_dir}/${bin_name} libCXXExample.dylib

    cd ${test_dir} && ./${bin_name}
}

function clean() {
    cd ${src_dir} && make clean
}

# prepare_cmake_toolchains
# update_cxx_library
update_protobuf
# init_go_mod
# build_go
# update_and_run_test
# clean