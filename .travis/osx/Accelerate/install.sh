#!/bin/bash
set -ex

export CGO_LDFLAGS="-framework Accelerate"
source ${TRAVIS_BUILD_DIR}/.travis/$TRAVIS_OS_NAME/install.sh
pushd cgo
go install -v -x
popd
