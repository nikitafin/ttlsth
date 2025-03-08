## ===--------------------------------------------------------------------=== ##
# TTLSTH - basic cxx project
## ===--------------------------------------------------------------------=== ##
cmake_minimum_required(VERSION 3.40)

## ===--------------------------------------------------------------------=== ##
# Common default options
set(CMAKE_CXX_STANDARD 26)
set(CMAKE_CXX_STANDARD_REQUIRED ON)
set(CMAKE_CXX_EXTENSIONS OFF)
set(CMAKE_CXX_SCAN_FOR_MODULES ON)
set(CMAKE_EXPORT_COMPILE_COMMANDS ON)

project(ttlsth
    VERSION 0.0.0
    LANGUAGES CXX
    DESCRIPTION "Basic cxx project"
    HOMEPAGE_URL "https://github.com/nikitafin/ttlsth"
)

## ===--------------------------------------------------------------------=== ##
# System Modules
include(CheckCXXCompilerFlag)
include(CheckLinkerFlag)
include(CMakePrintHelpers)

## ===--------------------------------------------------------------------=== ##
# User Modules (do not reorder)
include(cmake/PreventInSourceBuild.cmake)
include(cmake/ProjectVersion.cmake)
include(cmake/Detector.cmake)
include(cmake/Utilities.cmake)
include(cmake/HandleToolchain.cmake)
include(cmake/PrintSummary.cmake)

## ===--------------------------------------------------------------------=== ##
# Subprojects
add_subdirectory(projects)