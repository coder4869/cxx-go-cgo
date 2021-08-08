#!/bin/bash

script_dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd ${script_dir}

function copy_libs() {
    echo "copy_libs"
}

function relink_libs() {
    echo "relink_libs for ${1} : from @rpath/${2} to @executable_path/lib/${2}"
    install_name_tool -change "@rpath/${2}" "@executable_path/lib/${2}" ${1}
}


cmd=${1}
if [ ${cmd} == "copy" ]; then
    copy_libs
elif [ ${cmd} == "relink" ]; then
    app_name=${2}
    lib_name=${3}
    relink_libs ${app_name} ${lib_name}
fi