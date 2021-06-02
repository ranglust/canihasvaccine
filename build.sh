#!/bin/bash

myVersion=$(git tag | tail -1)
if [ -z "${myVersion}" ]; then
  myVersion="unknown"
fi


echo "Building canihasvaccine v${myVersion}"
GOOS=darwin go build -v -ldflags "-X  github.com/ranglust/canihasvaccine/cmd.VERSION=${myVersion}" -o "canihasvaccine-osx"
GOOS=windows go build -v -ldflags "-X  github.com/ranglust/canihasvaccine/cmd.VERSION=${myVersion}" -o "canihasvaccine-windows"
GOOS=linux go build -v -ldflags "-X  github.com/ranglust/canihasvaccine/cmd.VERSION=${myVersion}" -o "canihasvaccine-linux"

