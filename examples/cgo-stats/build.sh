#!/bin/bash

script_dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd ${script_dir}

src_dir=${script_dir}/src
build_dir=${script_dir}/build
bin_name=go_stats

function prepare() {
    # lib statgrab
    stats_dir=${script_dir}/../../tools/stats
    sh ${stats_dir}/statgrab.sh build
    # install
    cp -r ${stats_dir}/install/include  ${src_dir}/stats/include/
    cp -r ${stats_dir}/install/lib      ${src_dir}/stats/lib/
    # clean
    sh ${stats_dir}/statgrab.sh clean
}

function clean() {
    cd ${src_dir} && make clean
}


prepare

# go_cmd.sh init_go_mod ${src_dir} ${bin_name}
sh ${script_dir}/../../tools/scripts/go_cmd.sh init_go_mod ${src_dir} ${bin_name}
# # go_cmd.sh build_go ${src_dir} ${build_dir}
sh ${script_dir}/../../tools/scripts/go_cmd.sh build_go ${src_dir} ${build_dir}