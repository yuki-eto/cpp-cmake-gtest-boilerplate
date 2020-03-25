# cpp-cmake-gtest-boilerplate
CPP Shared Libs with GoogleTest by CMake

# cross platform build by Docker
install `Docker` and `Docker Compose`

```
$ docker-compose run --rm builder ./build_osx.sh
$ file build-osx/src/libsample_cpp_lib.dylib
build-osx/src/libsample_cpp_lib.dylib: Mach-O 64-bit dynamically linked shared library x86_64

$ docker-compose run --rm builder ./build_windows.sh
$ file build-windows/src/libsample_cpp_lib.dll
build-windows/src/libsample_cpp_lib.dll: PE32+ executable (DLL) (console) x86-64, for MS Windows
```
