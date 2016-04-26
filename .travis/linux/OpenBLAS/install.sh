#!/bin/bash
set -ex

# environment file for the docker container to use
ENV_FILE=$GOPATH/src/github.com/$TRAVIS_REPO_SLUG/.travis/$TRAVIS_OS_NAME/$BLAS_LIB/env.list
echo "TRAVIS_REPO_SLUG=$TRAVIS_REPO_SLUG" >> $ENV_FILE

# fetch repo
docker pull jonlawlor/gonum-docker-openblas

# install go, and get the gonum libs.
docker_travis_home="/root/gopath/src/github.com/$TRAVIS_REPO_SLUG"
docker run --name openblas \
  --env-file $ENV_FILE \
  -v $GOPATH/src:/root/gopath/src jonlawlor/gonum-docker-openblas \
  bash -c "eval \"\$(gimme $TRAVIS_GO_VERSION)\" \
    && source $docker_travis_home/.travis/$TRAVIS_OS_NAME/install.sh \
    && source $docker_travis_home/.travis/$TRAVIS_OS_NAME/$BLAS_LIB/docker-install.sh"

docker commit openblas jonlawlor/gonum-docker-openblas:temp
