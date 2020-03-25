#!/bin/bash

set -e

cd /workdir
BUILD_DIR=/workdir/build-windows
rm -rf $BUILD_DIR
mkdir $BUILD_DIR
cd $BUILD_DIR

export CROSS_TRIPLE=x86_64-w64-mingw32
crossbuild cmake ..
crossbuild make sample_cpp_lib

