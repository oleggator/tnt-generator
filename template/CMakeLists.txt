cmake_minimum_required(VERSION 3.15)
project(generator C)

set(CMAKE_C_STANDARD 11)

add_executable(generator
        funcs.c
        main.c
        models.c
        wrappers.c)

target_include_directories(generator PRIVATE /usr/local/include)
target_link_libraries (generator /usr/local/lib/libmsgpuck.a)
