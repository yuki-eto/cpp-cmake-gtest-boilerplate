#!/bin/bash

set -e

cd /workdir
BUILD_DIR=/workdir/build-osx
rm -rf $BUILD_DIR
mkdir $BUILD_DIR
cd $BUILD_DIR

export CROSS_TRIPLE=x86_64-apple-darwin
crossbuild cmake ..
sed -i -e "s/soname/install_name/g" src/CMakeFiles/sample_cpp_lib.dir/link.txt
crossbuild make sample_cpp_lib

