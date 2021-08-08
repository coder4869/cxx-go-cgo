set(LIB_NAME CXXPB)

set(SRC_DIR ${CMAKE_CURRENT_SOURCE_DIR})
file(GLOB_RECURSE src_files
    ${SRC_DIR}/*.cpp
    ${SRC_DIR}/*.cc
    ${SRC_DIR}/*.c++
    ${SRC_DIR}/*.c
    ${SRC_DIR}/*.hpp
    ${SRC_DIR}/*.h
    )

add_library(${LIB_NAME} SHARED ${src_files})
