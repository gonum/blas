#!/bin/bash
set -ex

# fetch and install ATLAS libs
sudo apt-get update -qq && sudo apt-get install -qq libatlas-base-dev

# fetch and install gonum/blas against ATLAS
export CGO_LDFLAGS="-L/usr/lib -lblas"
source ${TRAVIS_BUILD_DIR}/.travis/$TRAVIS_OS_NAME/install.sh
