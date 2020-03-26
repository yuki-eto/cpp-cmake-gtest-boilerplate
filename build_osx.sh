#!/bin/bash

set -e

cd /workdir
BUILD_DIR=/workdir/build-osx
rm -rf $BUILD_DIR
mkdir $BUILD_DIR
cd $BUILD_DIR

export CROSS_TRIPLE=x86_64-apple-darwin
crossbuild cmake -DCMAKE_SYSTEM_NAME=Darwin ..
crossbuild make calculator
