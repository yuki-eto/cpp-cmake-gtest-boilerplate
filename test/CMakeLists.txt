cmake_minimum_required(VERSION 3.15)

include(${PROJECT_SOURCE_DIR}/cmake/DownloadProject/DownloadProject.cmake)
download_project(PROJ googletest
                GIT_REPOSITORY https://github.com/google/googletest.git
                GIT_TAG master
                UPDATE_DISCONNECTED 1
)

add_subdirectory(${googletest_SOURCE_DIR} ${googletest_BINARY_DIR})

add_executable(TestLibrary test_calculator.cpp test_bridge.cpp)
target_link_libraries(TestLibrary calculator gtest_main)
include_directories(${PROJECT_SOURCE_DIR}/src)

include(GoogleTest)
gtest_add_tests(TARGET TestLibrary)
