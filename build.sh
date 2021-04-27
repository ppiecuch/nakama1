#!/bin/bash

set -e

echo_header() {
    if [[ "$TERM" =~ "xterm" ]]; then
        printf "\e[1;4m$1\e[0m\n"
    else
        printf "[$1]\n"
    fi
}

if [ -e $HOME/dev/go-bin ]; then
  ROOT=$HOME/dev/
elif [ -e /Volumes/Xta/dev/go-bin ]; then
  ROOT=/Volumes/Xta/dev
else
  echo_header "*** Cannot find Go instalation root path."
  exit 1
fi

export GO=$ROOT/go-bin
export GOBIN=$GO/bin
export GOPATH=$PWD
echo_header "*** GOPATH: $GOPATH"

PC=protoc-3.3.0-osx-x86_64
export PATH=/bin:/usr/bin:/usr/local/bin:$PWD/$PC/bin:$GOBIN:/opt/local/bin

if [ ! -d "src/nakama/vendor" ]; then
  (echo_header "*** Running: dep ensure" && cd src/nakama && dep ensure)
fi

(echo_header "*** Building: nakama" && cd src/nakama && make $*)
