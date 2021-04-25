#!/bin/bash

set -e

if [ -e $HOME/dev/go-bin ]; then
  ROOT=$HOME/dev/
elif [ -e /Volumes/Xta/dev/go-bin ]; then
  ROOT=/Volumes/Xta/dev
else
  echo "*** Cannot find Go instalation root path."
  exit 1
fi

export GO=$ROOT/go-bin
export GOBIN=$GO/bin
export GOPATH=$PWD
echo "*** GOPATH: $GOPATH"

PC=protoc-3.3.0-osx-x86_64
export PATH=/bin:/usr/bin:/usr/local/bin:$PWD/$PC/bin:$GOBIN:/opt/local/bin

(echo "*** dep ensure" && cd src/nakama && dep ensure)

(echo "*** build nakama" && cd src/nakama && make $*)
