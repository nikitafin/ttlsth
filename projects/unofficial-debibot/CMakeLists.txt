## ===--------------------------------------------------------------------=== ##
# Debibot: basic entry point
add_executable(debibot ${CMAKE_CURRENT_SOURCE_DIR}/cmd/Main.cpp)

set_target_properties(debibot
    PROPERTIES
    FOLDER unofficial-debibot
)