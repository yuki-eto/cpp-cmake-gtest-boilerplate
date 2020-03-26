#!/bin/bash

set -e

SRC_DIR=/workdir/src
cd $SRC_DIR

swig -go -c++ -cgo -intgosize 64 calculator.i
