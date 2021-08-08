#!/bin/bash

script_dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd ${script_dir}

function prepare_cmake_toolchains() {
    rm -rf ${script_dir}/cmake-toolchains
    # git clone https://github.com/coder4869/cmake-toolchains.git
    mkdir -p ${script_dir}/cmake-toolchains
    cp -rf ${script_dir}/../../../cmake-toolchains/* ${script_dir}/cmake-toolchains/
}

src_dir=${script_dir}/src
build_dir=${script_dir}/build
function build_cxx_with_cmake() {
    mkdir -p ${build_dir} && cd ${build_dir} && rm -rf *

    # run cmake
    cmake ${src_dir} \
        -G"Unix Makefiles" -DCMAKE_BUILD_TYPE=Release -H${src_dir} -B${build_dir}
    make -j
}

output_dir=${script_dir}/output
function update_output() {
    rm -rf ${output_dir}/* && mkdir -p ${output_dir}/include/ ${output_dir}/lib/

    # copy headers to output
    cd ${src_dir} && find . \( -iname "*.h" -o -iname "*.hpp" \) |xargs tar czvf include.tgz
    tar -zxvf include.tgz -C ${output_dir}/include/ && rm -rf include.tgz

    # copy libs to output
    find ${build_dir} -name "lib*" -exec cp {} ${output_dir}/lib/ \;
}

function clean() {
    rm -rf ${build_dir} # ${output_dir}
}

# prepare_cmake_toolchains
build_cxx_with_cmake
update_output
clean