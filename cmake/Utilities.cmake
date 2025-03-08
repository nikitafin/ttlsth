## ===--------------------------------------------------------------------=== ##
# Various utilities
## ===--------------------------------------------------------------------=== ##
include_guard(GLOBAL)
message(STATUS "Load user CMake module: ${CMAKE_CURRENT_LIST_FILE}")

function(check_and_append_compiler_flag flag list)
    # Generate a variable name from the flag
    set(flag_var Compiler:${flag})

    # Check if the compiler supports the flag
    check_cxx_compiler_flag(${flag} ${flag_var})

    if(${flag_var})
        # Check if the flag is not already in the list
        list(FIND ${list} ${flag} flag_index)

        if(flag_index EQUAL -1)
            # Append the flag to the list
            list(APPEND ${list} ${flag})
            set(${list} ${${list}} PARENT_SCOPE)
        endif()
    endif()
endfunction()

function(check_and_append_linker_flag flag list)
    # Generate a variable name from the flag
    set(flag_var Linker:${flag})

    # Check if the compiler supports the flag
    check_linker_flag(CXX ${flag} ${flag_var})

    if(${flag_var})
        # Check if the flag is not already in the list
        list(FIND ${list} ${flag} flag_index)

        if(flag_index EQUAL -1)
            # Append the flag to the list
            list(APPEND ${list} ${flag})
            set(${list} ${${list}} PARENT_SCOPE)
        endif()
    endif()
endfunction()

function(append_unique list_name value)
    if(NOT ${list_name})
        set(${list_name} ${value} PARENT_SCOPE)
    else()
        list(FIND ${list_name} ${value} index)
        if(index EQUAL -1)
            list(APPEND ${list_name} ${value})
            set(${list_name} ${${list_name}} PARENT_SCOPE)
        endif()
    endif()
endfunction()
