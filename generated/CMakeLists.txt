project(myproject C)
cmake_minimum_required(VERSION 3.14)

set(CMAKE_C_STANDARD 11)
include_directories(/usr/local/include)

add_library(cproc SHARED models.c models.h)
target_link_libraries(cproc msgpuck)
