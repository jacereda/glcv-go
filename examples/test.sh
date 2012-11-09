#!/bin/sh
export LD_LIBRARY_PATH=../cv/glcv/b
export DYLD_LIBRARY_PATH=../cv/glcv/b
go build test.go
./test

