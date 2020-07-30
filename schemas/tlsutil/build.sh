#!/bin/bash

# gets script dir
SRCDIR=$(dirname $(readlink -f "$0"))

# sets path for tools
GOPATH="${GOPATH:-$HOME/go}"
PATH=$GOPATH/bin:$PATH

protoc -I $GOPATH/src -I . $SRCDIR/common.proto --go_out=plugins=grpc:$GOPATH/src
protoc -I $GOPATH/src -I . $SRCDIR/archive.proto --go_out=plugins=grpc:$GOPATH/src
protoc -I $GOPATH/src -I . $SRCDIR/classify.proto --go_out=plugins=grpc:$GOPATH/src
protoc -I $GOPATH/src -I . $SRCDIR/analyze.proto --go_out=plugins=grpc:$GOPATH/src
protoc -I $GOPATH/src -I . $SRCDIR/notary.proto --go_out=plugins=grpc:$GOPATH/src
