#!/usr/bin/env bash

source ./common.sh

function start() {
    cd ${servicePath} && \
    go run admin.go -f etc/admin-api.yaml
}

start