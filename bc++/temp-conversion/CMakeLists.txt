cmake_minimum_required(VERSION 3.7)
project(temp_conversion)

set(CMAKE_CXX_STANDARD 11)

set(SOURCE_FILES src/main.cpp src/temp_conversion.hpp test/celcius_to_farenheit_test.cpp test/celcius_to_farenheit_test.h)
add_executable(temp_conversion ${SOURCE_FILES})