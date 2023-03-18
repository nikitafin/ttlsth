macro(set_msvc_compiler_flags)
    add_compile_options(
            /W4
            /WX # Treats warnings as errors
            /wd4710 # Suppresses warning C4710: function 'xyz' not inlined
            /wd4820 # Suppresses warning C4820: 'bytes' bytes padding added after data member
            /wd5050 # Disable module waning
            /EHsc # Enable exceptions
            /experimental:module # cxx20 module support
            /std:c++latest # cxx standard
            /permissive- # ISO cxx
            /fp:precise # precision floating point calculation
    )

    if (CMAKE_BUILD_TYPE MATCHES Debug)
        add_compile_options(
                /Od # Disable optimization to make debugging easier
                /Zi # Generates debugging information that can be used by a debugger
                /sdl # Runtime security checks
                /RTC1 # Enables run-time error checks
        )

        add_compile_definitions("_DEBUG")
    else ()
        add_compile_options(
                /O2 # Enable maximum optimization
                /Oi # Enables intrinsic functions
                /GL # Enables whole program optimization
        )
    endif ()
endmacro()