add_executable(rt main.cc common.ixx)

add_executable(common_test common_test.cc common.ixx)
target_link_libraries(common_test PRIVATE shared::gtest_all)

enable_testing()
add_test(NAME common_test COMMAND common_test WORKING_DIRECTORY ${CMAKE_SOURCE_DIR})