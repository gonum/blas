#!/bin/bash
set -ex

go generate github.com/gonum/blas/cgo
go generate github.com/gonum/blas/native
if [ -n "$(git diff ${TRAVIS_BUILD_DIR}/cgo)" ]; then
  echo "files changed in /cgo"
  exit 1
fi

if [ -n "$(git diff ${TRAVIS_BUILD_DIR}/native)" ]; then
  echo "files changed in /native"
  exit 1
fi
