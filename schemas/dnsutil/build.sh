#!/bin/bash

# gets script dir
SRCDIR=$(dirname $(readlink -f "$0"))

# sets path for tools
GOPATH="${GOPATH:-$HOME/go}"
PATH=$GOPATH/bin:$PATH

protoc -I $GOPATH/src -I . $SRCDIR/resolv.proto --go_out=plugins=grpc:$GOPATH/src
protoc -I $GOPATH/src -I . $SRCDIR/archive.proto --go_out=plugins=grpc:$GOPATH/src
protoc -I $GOPATH/src -I . $SRCDIR/finder.proto --go_out=plugins=grpc:$GOPATH/src
