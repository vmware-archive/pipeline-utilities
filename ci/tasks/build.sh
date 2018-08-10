#!/bin/bash -e

mkdir ~/.ssh/ && touch ~/.ssh/known_hosts
ssh-keyscan github.com >>~/.ssh/known_hosts

export GOPATH=$PWD/go
export PATH=$GOPATH/bin:$PATH
OUTPUT_DIR=$PWD/compiled-output
SOURCE_DIR=$PWD/source

cp source/Dockerfile ${OUTPUT_DIR}/.

go get -u github.com/golang/dep/cmd/dep
go get github.com/xchapter7x/versioning

cd ${SOURCE_DIR}
if [ -d ".git" ]; then
  DRAFT_VERSION=`versioning bump_patch`-`git rev-parse HEAD`
else
  DRAFT_VERSION="v0.0.0-local"
fi
echo "next version should be: ${DRAFT_VERSION}"

WORKING_DIR=$GOPATH/src/github.com/pivotalservices/pipeline-utilities
mkdir -p ${WORKING_DIR}
cp -R ${SOURCE_DIR}/* ${WORKING_DIR}/.
cd ${WORKING_DIR}
dep ensure
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${OUTPUT_DIR}/pipeline-utilities-linux -ldflags "-X github.com/pivotalservices/pipeline-utilities/commands.VERSION=${DRAFT_VERSION}" cmd/pipeline-utilities/main.go
GOOS=darwin GOARCH=amd64 go build -o ${OUTPUT_DIR}/pipeline-utilities-osx -ldflags "-X github.com/pivotalservices/pipeline-utilities/commands.VERSION=${DRAFT_VERSION}" cmd/pipeline-utilities/main.go
GOOS=windows GOARCH=amd64 go build -o ${OUTPUT_DIR}/pipeline-utilities.exe -ldflags "-X github.com/pivotalservices/pipeline-utilities/commands.VERSION=${DRAFT_VERSION}" cmd/pipeline-utilities/main.go

echo ${DRAFT_VERSION} > ${OUTPUT_DIR}/name
echo ${DRAFT_VERSION} > ${OUTPUT_DIR}/tag
