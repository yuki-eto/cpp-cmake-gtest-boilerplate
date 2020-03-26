#!/bin/bash

set -e

BUILD_DIR=/usr/local/build
rm -rf $BUILD_DIR
mkdir $BUILD_DIR
cd $BUILD_DIR

cmake /workdir
make
valgrind --leak-check=full --leak-resolution=med --track-origins=yes --vgdb=no ./test/TestLibrary
