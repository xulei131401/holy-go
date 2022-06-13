#!/usr/bin/env bash

servicePath="/Users/xulei/jungle/githubproject/holy-go/service/api/admin"

function start() {
    cd ${servicePath} && \
    go run admin.go -f etc/admin-api.yaml
}

start