#!/bin/bash

script_dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd ${script_dir}

function prepare_cmake_toolchains() {
    rm -rf ${script_dir}/cmake-toolchains
    # git clone https://github.com/coder4869/cmake-toolchains.git
    mkdir -p ${script_dir}/cmake-toolchains
    cp -rf ${script_dir}/../../../cmake-toolchains/* ${script_dir}/cmake-toolchains/
}

function update_cxx_library() {
    sh ${script_dir}/../cgo-cxx/build.sh
    cp -r ${script_dir}/../cgo-cxx/output/ ${script_dir}/library
}

################# build_go #################
src_dir=${script_dir}/src
build_dir=${script_dir}/build
bin_name=cgo_base

test_dir=${script_dir}/test/
function update_and_run_test() {
    mkdir -p ${test_dir} && rm -rf ${test_dir}/* 

    mv ${build_dir}/bin/* ${test_dir}
    sh ${script_dir}/scripts/lib_process.sh relink ${test_dir}/${bin_name} libCXXBase.dylib

    cd ${test_dir} && ./${bin_name}
}

function clean() {
    cd ${src_dir} && make clean
}

# prepare_cmake_toolchains
update_cxx_library

# go_cmd.sh init_go_mod ${src_dir} ${bin_name}
sh ${script_dir}/../../tools/scripts/go_cmd.sh init_go_mod ${src_dir} ${bin_name}
# go_cmd.sh build_go ${src_dir} ${build_dir}
sh ${script_dir}/../../tools/scripts/go_cmd.sh build_go ${src_dir} ${build_dir}

update_and_run_test
clean