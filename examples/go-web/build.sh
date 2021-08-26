#!/bin/bash

script_dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd ${script_dir}

src_dir=${script_dir}/src
build_dir=${script_dir}/build
bin_name=goweb


# go_cmd.sh init_go_mod ${src_dir} ${bin_name}
sh ${script_dir}/../../tools/scripts/go_cmd.sh init_go_mod ${src_dir} ${bin_name}
# # go_cmd.sh build_go ${src_dir} ${build_dir}
sh ${script_dir}/../../tools/scripts/go_cmd.sh build_go ${src_dir} ${build_dir}