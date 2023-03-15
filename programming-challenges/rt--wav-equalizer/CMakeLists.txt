add_library(common)
target_link_libraries(common)

target_sources(common
    PUBLIC
    FILE_SET cxx_modules TYPE CXX_MODULES FILES
    common.ixx
)

add_executable(rt main.cc) #common.ixx)
#get_target_property(cpl rt COMPILE_OPTIONS)
#message(${cpl})
target_link_libraries(rt common)