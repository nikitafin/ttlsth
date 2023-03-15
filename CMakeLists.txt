cmake_minimum_required(VERSION 3.25)

project(ttlsth)
set(CMAKE_CXX_STANDARD 23)
set(CMAKE_C_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)
set(CMAKE_C_STANDARD_REQUIRED ON)
set(CMAKE_POSITION_INDEPENDENT_CODE ON)
set(CMAKE_MSVC_RUNTIME_LIBRARY "MultiThreaded$<$<CONFIG:Debug>:Debug>")
set(CMAKE_EXPERIMENTAL_CXX_MODULE_DYNDEP 1)

# turn on the experimental API
set(CMAKE_EXPERIMENTAL_CXX_MODULE_CMAKE_API "3c375311-a3c9-4396-a187-3227ef642046")

# Collect all of the .cmake files in the cmake directory
file(GLOB CMAKE_FILES "${CMAKE_CURRENT_LIST_DIR}/cmake/*.cmake")

# Include all of the collected cmake files
foreach(CMAKE_FILE ${CMAKE_FILES})
    include(${CMAKE_FILE})
endforeach()

if(CMAKE_C_COMPILER_ID MATCHES "MSVC")
    set(CMAKE_EXPERIMENTAL_CXX_MODULE_DYNDEP 1)
    string(CONCAT CMAKE_EXPERIMENTAL_CXX_SCANDEP_SOURCE
        "<CMAKE_CXX_COMPILER> <DEFINES> <INCLUDES> <FLAGS> <SOURCE> -nologo -TP"
        " -showIncludes"
        " -scanDependencies <DYNDEP_FILE>"
        " -Fo<OBJECT>")
    set(CMAKE_EXPERIMENTAL_CXX_SCANDEP_DEPFILE_FORMAT "msvc")
    set(CMAKE_EXPERIMENTAL_CXX_MODULE_MAP_FORMAT "msvc")
    set(CMAKE_EXPERIMENTAL_CXX_MODULE_MAP_FLAG "@<MODULE_MAP_FILE>")
    set_msvc_compiler_flags()
elseif(CMAKE_C_COMPILER_ID MATCHES "GNU|MinGW")
    set_gcc_compiler_flags()
elseif(CMAKE_C_COMPILER_ID MATCHES "Clang")
    set_clang_compiler_flags()
endif()

enable_testing()

# add_subdirectory(shared/googletest)
add_subdirectory(programming-challenges)
