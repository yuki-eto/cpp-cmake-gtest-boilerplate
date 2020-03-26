# cpp-cmake-gtest-boilerplate
CPP Shared Libs with GoogleTest by CMake

# build and test
```
mkdir build-cmake
cd build-cmake
cmake ..

# build all
make

# test
make test
```

# cross platform build by Docker
install `Docker` and `Docker Compose`

```
$ docker-compose run --rm builder ./build_osx.sh
$ file build-osx/src/libcalculator.dylib
build-osx/src/libcalculator.dylib: Mach-O 64-bit dynamically linked shared library x86_64

$ docker-compose run --rm builder ./build_windows.sh
$ file build-windows/src/libcalculator.dll
build-windows/src/libcalculator.dll: PE32+ executable (DLL) (console) x86-64, for MS Windows
```
