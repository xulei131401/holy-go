#!/usr/bin/env bash

projectPath="/Users/xulei/jungle/githubproject/holy-go"
apiPath="/Users/xulei/jungle/githubproject/holy-go/script/api/admin"
servicePath="/Users/xulei/jungle/githubproject/holy-go/service/api/admin"

function rebuild() {
    cd ${projectPath} && \
    goctl api go -api ${apiPath}/admin.api -dir ${servicePath}
}

rebuild