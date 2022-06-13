#!/usr/bin/env bash

projectPath="/Users/xulei/jungle/githubproject/holy-go"
modelPath="/Users/xulei/jungle/githubproject/holy-go/service/api/admin/internal/model"
sqlPath="/Users/xulei/jungle/githubproject/holy-go/sql"

function rebuild() {
  #-c 表示生成带缓存的代码
  cd ${projectPath} &&
#   goctl model mysql ddl -src="${sqlPath}/*.sql" -dir=${modelPath}
  goctl model mysql datasource -url="root:xl123456?@tcp(47.105.50.31:3306)/holy-admin" -table="*" -dir=${modelPath}
}

rebuild
