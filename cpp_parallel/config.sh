#!/bin/bash

rm -rf build
CC=clang++ CXX=clang++ CFLAGS=-stdlib=libc++ LDFLAGS=-lc++abi meson build
