#!/bin/bash

script_dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd ${script_dir}

pkg=libstatgrab-0.91
function prepare() {
    if [ ! -f "${pkg}.tar.gz" ]; then
        curl -o ${pkg}.tar.gz http://www.mirrorservice.org/pub/i-scream/libstatgrab/${pkg}.tar.gz
    fi
}

function build() {
    tar -zvxf ${pkg}.tar.gz
    cd ${pkg}
    ./configure --prefix=${script_dir}/install
    make install
    cd -
}

function clean() {
    rm -rf ${pkg} install # ${pkg}.tar.gz
}

if [ "$1" == "build" ]; then
    prepare
    build
elif [ "$1" == "clean" ]; then
    clean
fi
