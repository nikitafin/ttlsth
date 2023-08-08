add_library(
        pacardolib STATIC
        pacardo.h

        src/pacardolib.cc
        src/pacardolib.h
)
target_include_directories(pacardolib PUBLIC ${CMAKE_CURRENT_SOURCE_DIR}/)

enable_testing()

add_executable(pacardolib_test src/pacardolib.t.cc)
target_link_libraries(pacardolib_test PRIVATE pacardolib shared::gtest_all)
add_test(NAME pacardolib_test COMMAND pacardolib_test WORKING_DIRECTORY ${CMAKE_SOURCE_DIR})