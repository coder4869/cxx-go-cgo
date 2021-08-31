cmake_minimum_required(VERSION 3.17)

set(SDK_NAME CXXExample)
project(${SDK_NAME})

# FOR SHARED LIBRARY, OTHERWISE SUPPORTS STATIC ONLY
SET_PROPERTY(GLOBAL PROPERTY USE_FOLDERS ON)
SET_PROPERTY(GLOBAL PROPERTY TARGET_SUPPORTS_SHARED_LIBS TRUE)

# wasm
# set(WASM_OUTPUT_NAME lib${SDK_NAME})
# include(cmake-toolchains/emscripten-options.cmake)

set(CMAKE_CXX_FLAGS "-std=c++14 -O2")

add_subdirectory(cBase)
add_subdirectory(cxxBase)
add_subdirectory(pb)