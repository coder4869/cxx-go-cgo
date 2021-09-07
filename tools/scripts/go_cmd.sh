#!/bin/bash

################# init go mod #################
# call：init_go_mod ${src_dir} ${bin_name}
function init_go_mod() {
    src_dir=${1}
    bin_name=${2}

    cd ${src_dir}
    if [ ! -f "go.mod" ]; then
        go mod init ${bin_name}
    fi

    if [ -f "go.mod" ]; then
       go mod tidy
    fi
}


################# build go #################
# call：build_go ${src_dir} ${build_dir}
function build_go() {
    src_dir=${1}
    build_dir=${2}

    mkdir -p ${build_dir} && rm -rf ${build_dir}/* 

    make -C ${src_dir} # -C go to target dir and run Makefile
}

if [ "$1" == "init_go_mod" ]; then
    init_go_mod ${2} ${3}
elif [ "$1" == "build_go" ]; then
    build_go ${2} ${3}
fi