# pgp_add_shared cmake_folder[ base_folder1[, ...base_folderN]]
# Copied from https://github.com/ClickHouse/ClickHouse/blob/f1df91cd0998ee7ba368972855bf16c84d920a9f/contrib/CMakeLists.txt#LL25
function(pgp_add_shared cmake_folder)
    if (ARGN)
        set(base_folders ${ARGN})
    else ()
        set(base_folders ${cmake_folder})
    endif ()

    foreach (base_folder ${base_folders})
        # some typos in the code
        if (NOT IS_DIRECTORY "${CMAKE_CURRENT_SOURCE_DIR}/${base_folder}")
            message(FATAL_ERROR "No such base folder '${base_folder}' (for '${cmake_folder}' cmake folder). Typo in the base folder name?")
        endif ()

        file(GLOB contrib_files "${base_folder}/*")
        if (NOT contrib_files)
            # Checking out *all* submodules takes > 5 min. Therefore, the smoke build ("FastTest") in CI initializes only the set of
            # submodules minimally needed for a build and we cannot assume here that all submodules are populated.
            message(STATUS "submodule ${base_folder} is missing or empty. to fix try run:")
            message(STATUS "    git submodule update --init")
            return()
        endif ()
    endforeach ()

    message(STATUS "[Shared]: Adding module ${base_folders} (configuring with ${cmake_folder})")
    add_subdirectory(${cmake_folder})
endfunction()