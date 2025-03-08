## ===--------------------------------------------------------------------=== ##
# Prevent build
## ===--------------------------------------------------------------------=== ##

cmake_path(SET PROJECT_SRC_DIR "${CMAKE_SOURCE_DIR}" NORMALIZE)
cmake_path(SET PROJECT_BIN_DIR "${CMAKE_BINARY_DIR}" NORMALIZE)

if(PROJECT_SRC_DIR STREQUAL PROJECT_BIN_DIR)
    message("## ===--------------------------------------------------------------------=== ##")
    message("
In-source builds are disabled to prevent cluttering the source directory.

Please create a separate build directory and run CMake from there. For example:

    mkdir -p build
    cd build
    cmake ..

Alternatively, you can specify an out-of-source build directory directly:

    cmake -B build -S .
")
    message("## ===--------------------------------------------------------------------=== ##")
    message(FATAL_ERROR "Quitting configuration")
endif()
