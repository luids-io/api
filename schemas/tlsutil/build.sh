#!/bin/bash

# gets script dir
SRCDIR=$(dirname $(readlink -f "$0"))

# sets path for tools
GOPATH="${GOPATH:-$HOME/go}"
PATH=$GOPATH/bin:$PATH

protoc -I $GOPATH/src -I . common.proto --go_out=plugins=grpc:$GOPATH/src
protoc -I $GOPATH/src -I . archive.proto --go_out=plugins=grpc:$GOPATH/src
protoc -I $GOPATH/src -I . classify.proto --go_out=plugins=grpc:$GOPATH/src
protoc -I $GOPATH/src -I . analyze.proto --go_out=plugins=grpc:$GOPATH/src
protoc -I $GOPATH/src -I . notary.proto --go_out=plugins=grpc:$GOPATH/src
