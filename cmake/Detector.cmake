## ===--------------------------------------------------------------------=== ##
# Detector - detect configuration
## ===--------------------------------------------------------------------=== ##
include_guard(GLOBAL)
message(STATUS "Load user CMake module: ${CMAKE_CURRENT_LIST_FILE}")

## ===--------------------------------------------------------------------=== ##
# Supported compilers
set(PROJECT_COMPILER_CLANG OFF)
set(PROJECT_COMPILER_GCC OFF)
set(PROJECT_COMPILER_CL OFF)
set(PROJECT_COMPILER_CLANG_CL OFF)
set(PROJECT_COMPILER_MINGW_CLANG OFF)
set(PROJECT_COMPILER_MINGW_GCC OFF)

# TODO(nikitafin): support more compiler

## ===--------------------------------------------------------------------=== ##
# Detect compiler
if(WIN32)
    if(CMAKE_CXX_COMPILER_ID STREQUAL "MSVC")
        set(PROJECT_COMPILER_CL ON)
    else()
        message(FATAL_ERROR "Unsupported compiler: ${CMAKE_CXX_COMPILER_ID} (${CMAKE_CXX_COMPILER})")
    endif()
elseif(LINUX)
    if(CMAKE_CXX_COMPILER_ID STREQUAL "Clang")
        set(PROJECT_COMPILER_CLANG ON)
    elseif(CMAKE_CXX_COMPILER_ID STREQUAL "GNU")
        set(PROJECT_COMPILER_GCC ON)
    else()
        message(FATAL_ERROR "Unsupported compiler: ${CMAKE_CXX_COMPILER_ID} (${CMAKE_CXX_COMPILER})")
    endif()
else()
endif()

# TODO(nikitafin): add more sanitizer support
## ===--------------------------------------------------------------------=== ##
# Sanitizer Support
set(PROJECT_SANITIZER_ADDRESS_SUPPORTED OFF)

if(PROJECT_COMPILER_CL)
    set(PROJECT_SANITIZER_ADDRESS_SUPPORTED ON)
endif()

set(PROJECT_SANITIZER_USED OFF)
set(PROJECT_SANITIZER_LIST)

if(PROJECT_SANITIZER_ADDRESS_SUPPORTED AND PROJECT_SANITIZER_ADDRESS_ENABLED)
    set(PROJECT_SANITIZER_USED ON)
    list(APPEND PROJECT_SANITIZER_LIST "address")
endif()

## ===--------------------------------------------------------------------=== ##
# 
set(PROJECT_BUILD_TYPE "DEBUG")
string(TOUPPER ${CMAKE_BUILD_TYPE} PROJECT_BUILD_TYPE)

## ===--------------------------------------------------------------------=== ##
# Toolchain default options
set(PROJECT_DEFAULT_COMPILE_OPTION ${CMAKE_CXX_FLAGS} ${CMAKE_CXX_FLAGS_${PROJECT_BUILD_TYPE}})
separate_arguments(PROJECT_DEFAULT_COMPILE_OPTION)

set(PROJECT_DEFAULT_LINK_OPTIONS ${CMAKE_EXE_LINKER_FLAGS} ${CMAKE_EXE_LINKER_FLAGS_${PROJECT_BUILD_TYPE}})
separate_arguments(PROJECT_DEFAULT_LINK_OPTIONS)

