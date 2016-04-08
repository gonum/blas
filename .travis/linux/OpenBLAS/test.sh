#!/bin/bash
set -ex

ENV_FILE=$GOPATH/src/github.com/$TRAVIS_REPO_SLUG/.travis/$TRAVIS_OS_NAME/$BLAS_LIB/env.list

# run the docker-test script from within the container
docker run \
  --env-file $ENV_FILE \
  -v $GOPATH/src:/root/gopath/src jonlawlor/gonum-docker-openblas:temp \
  bash -c "eval \"\$(gimme $TRAVIS_GO_VERSION)\" && \
    cd /root/gopath/src/github.com/$TRAVIS_REPO_SLUG/ && \
    source /root/gopath/src/github.com/$TRAVIS_REPO_SLUG/.travis/$TRAVIS_OS_NAME/$BLAS_LIB/docker-test.sh"
