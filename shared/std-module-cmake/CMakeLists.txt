if (MSVC_MODULE)
    add_library(std_module OBJECT "${MSVC_MODULE}/modules/std.ixx")
endif ()
