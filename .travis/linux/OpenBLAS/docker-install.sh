#!/bin/bash
set -ex

cd $GOPATH/src/github.com/$TRAVIS_REPO_SLUG/cgo
go install -v -x
