macro(set_clang_compiler_flags)
    add_compile_options(
        -Wall # Enables all warnings
        -Wextra # Enables extra warnings not enabled by -Wall
        -Werror # Treats warnings as errors
    )

    if(CMAKE_BUILD_TYPE MATCHES Debug)
        add_compile_options(
            -O0 # Disable optimization to make debugging easier
            -g # Generates debugging information that can be used by a debugger
        )
    elseif(CMAKE_BUILD_TYPE MATCHES Release)
        add_compile_options(
            -O3 # Enable maximum optimization
        )
    elseif(CMAKE_BUILD_TYPE MATCHES RelWithDebInfo)
        add_compile_options(
            -O2 # Enable optimization
            -g # Generates debugging information that can be used by a debugger
        )
    endif()
endmacro()
