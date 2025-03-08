include_guard(GLOBAL)

message(STATUS "Load user CMake module: ${CMAKE_CURRENT_LIST_FILE}")

find_package(Git QUIET)
if(NOT GIT_FOUND)
    message(WARNING "Cannot find Git package. Cannot fetch PROJECT_VERSION")
    return()
endif()

# Get the latest git tag
execute_process(
    COMMAND Git::Git describe --tags
    WORKING_DIRECTORY ${CMAKE_SOURCE_DIR}
    OUTPUT_VARIABLE GIT_TAG
    OUTPUT_STRIP_TRAILING_WHITESPACE
)

if(NOT GIT_TAG)
    set(PROJECT_VERSION_MAJOR 0)
    set(PROJECT_VERSION_MINOR 0)
    set(PROJECT_VERSION_PATH 0)
else()
    string(REPLACE "v" "" GIT_TAG ${GIT_TAG})

    string(REPLACE "." ";" VERSION_LIST ${GIT_TAG})

    list(GET VERSION_LIST 0 PROJECT_VERSION_MAJOR)
    list(GET VERSION_LIST 1 PROJECT_VERSION_MINOR)
    list(GET VERSION_LIST 2 PROJECT_VERSION_PATCH)
endif()

set(PROJECT_VERSION ${PROJECT_VERSION_MAJOR}.${PROJECT_VERSION_MINOR}.${PROJECT_VERSION_PATH})


if(PROJECT_CMAKE_DEBUG_LOG_ENABLED)
    cmake_print_variables(PROJECT_VERSION)
endif()
