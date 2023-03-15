macro(set_msvc_compiler_flags)
    add_compile_options(
        /W4 # Enables a high level of warning messages
        /WX # Treats warnings as errors
        /wd4710 # Suppresses warning C4710: function 'xyz' not inlined
        /wd4820 # Suppresses warning C4820: 'bytes' bytes padding added after data member
        /EHsc
        /experimental:module
        /std:c++latest
        /translateInclude
    )

    if(CMAKE_BUILD_TYPE MATCHES Debug)
        add_compile_options(
            /Od # Disable optimization to make debugging easier
            /Zi # Generates debugging information that can be used by a debugger
            /sdl
            /RTC1 # Enables run-time error checks
            /MTd
        )
    else()
        add_compile_options(
            /O2 # Enable maximum optimization
            /Oi # Enables intrinsic functions
            /GL # Enables whole program optimization
            /MT
        )
    endif()
endmacro()