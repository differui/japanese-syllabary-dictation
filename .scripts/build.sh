#!/bin/bash

function build() {
    local NAME=jsd
    $(go env GOPATH)/bin/go-bindata ./assets
    go build -o $NAME
}

build