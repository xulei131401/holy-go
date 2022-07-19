#!/usr/bin/env bash

source ./common.sh

function rebuild() {
    cd ${projectPath} && \
    goctl api go --api ${apiPath}/admin.api --home $GOCTL_TEMPLATE --dir ${servicePath}
    # 项目与goctl不一致的行为可以自定义
    # go run 项目目录/holy-go/cmd/main.go
}

rebuild