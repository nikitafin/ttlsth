## ===--------------------------------------------------------------------=== ##
# Handle various toolchains (MOST UGLY PART)
## ===--------------------------------------------------------------------=== ##
include_guard(GLOBAL)
message(STATUS "Load user CMake module: ${CMAKE_CURRENT_LIST_FILE}")

## ===--------------------------------------------------------------------=== ##
# Reset default options
set(CMAKE_CXX_FLAGS "")
set(CMAKE_CXX_FLAGS_${PROJECT_BUILD_TYPE} "")
set(CMAKE_CXX_STANDARD_LIBRARIES "")
set(CMAKE_EXE_LINKER_FLAGS "")
set(CMAKE_EXE_LINKER_FLAGS_${PROJECT_BUILD_TYPE} "")

## ===--------------------------------------------------------------------=== ##
# Default options
foreach(opt ${PROJECT_DEFAULT_COMPILE_OPTION})
    append_unique(PROJECT_COMPILE_OPTIONS ${opt})
endforeach()

if(NOT PROJECT_COMPILE_OPTIONS)
    set(PROJECT_COMPILE_OPTIONS)
endif()

foreach(opt ${PROJECT_DEFAULT_LINK_OPTIONS})
    append_unique(PROJECT_LINK_OPTIONS ${opt})
endforeach()

if(NOT PROJECT_LINK_OPTIONS)
    set(PROJECT_LINK_OPTIONS)
endif()

## ===--------------------------------------------------------------------=== ##
# сl compiler common options
if(PROJECT_COMPILER_CL)
    append_unique(PROJECT_COMPILE_OPTIONS "/permissive-")
endif()

## ===--------------------------------------------------------------------=== ##
# sanitizers
if(PROJECT_SANITIZER_USED)
    # cl
    if(PROJECT_COMPILER_CL)
        # limitations:
        #    1. The /RTC options are incompatible with AddressSanitizer and should be disabled.
        #    2. Incremental linking is unsupported, and should be disabled.
        #    3. Edit-and-Continue is unsupported, and should be disabled.
        #    4. Coroutines are incompatible with AddressSanitizer, and resumable functions are exempt from instrumentation.
        #    5. OpenMP is unsupported, and should be disabled.
        #    6. Managed C++ is unsupported, and should be disabled.
        #    7. C++ AMP is unsupported, and should be disabled.
        #    8. Universal Windows Platform (UWP) applications are unsupported.
        #    9. Special case list files are unsupported.
        #   10. MFC using the optional alloc_dealloc_mismatch runtime option is unsupported.
        append_unique(PROJECT_COMPILE_OPTIONS "/fsanitize=address")
        if(NOT MSVC_IDE)
            append_unique(PROJECT_COMPILE_OPTIONS "/fno-sanitize-address-vcasan-lib")
        endif()

        # NOTE: By default (3.31.3) cmake automatically add -Zi into compiler flags
        if(NOT CMAKE_BUILD_TYPE STREQUAL "Debug" AND NOT CMAKE_BUILD_TYPE STREQUAL "RelWithDebugInfo")
            append_unique(PROJECT_COMPILE_OPTIONS "/Zi")
        endif()

        list(FILTER PROJECT_COMPILE_OPTIONS EXCLUDE REGEX "RTC") # 1
        list(FILTER PROJECT_LINK_OPTIONS EXCLUDE REGEX "INCREMENTAL") # 2

        append_unique(PROJECT_LINK_OPTIONS "/INCREMENTAL:NO")
    endif()
endif()

## ===--------------------------------------------------------------------=== ##
# Common defines
if(WIN32 OR MINGW)
    append_unique(PROJECT_COMPILE_DEFINITIONS "_WIN32")
    append_unique(PROJECT_COMPILE_DEFINITIONS "_USE_MATH_DEFINES")
    append_unique(PROJECT_COMPILE_DEFINITIONS "NOMINIMAX")
    append_unique(PROJECT_COMPILE_DEFINITIONS "NOGDI")
endif()

## ===--------------------------------------------------------------------=== ##
# Apply

if(PROJECT_COMPILE_OPTIONS)
    add_compile_options(${PROJECT_COMPILE_OPTIONS})
endif()

if(PROJECT_COMPILE_DEFINITIONS)
    add_compile_definitions(${PROJECT_COMPILE_DEFINITIONS})
endif()

if(PROJECT_LINK_OPTIONS)
    add_link_options(${PROJECT_LINK_OPTIONS})
endif()
