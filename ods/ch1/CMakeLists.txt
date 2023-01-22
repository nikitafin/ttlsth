add_executable(ex1.1 ex1.1.cc)

add_executable(ex1.1_test ex1.1_test.cc)
target_link_libraries(ex1.1_test PUBLIC gtest_main)

include(GoogleTest)
gtest_discover_tests(ex1.1_test)