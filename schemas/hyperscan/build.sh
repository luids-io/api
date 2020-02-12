#!/bin/bash

# gets script dir
SRCDIR=$(dirname $(readlink -f "$0"))

# sets path for tools
GOPATH="${GOPATH:-$HOME/go}"
PATH=$GOPATH/bin:$PATH

protoc -I $GOPATH/src -I . hyperscan.proto --go_out=plugins=grpc:$GOPATH/src
