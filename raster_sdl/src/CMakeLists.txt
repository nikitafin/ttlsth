project(main)

add_executable(${PROJECT_NAME} WIN32 main.cc)
target_link_libraries(${PROJECT_NAME} PRIVATE SDL2-static SDL2main)