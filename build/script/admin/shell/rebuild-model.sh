#!/usr/bin/env bash

source ./common.sh

function rebuild() {
  #-c 表示生成带缓存的代码
  #   goctl model mysql ddl -src="${sqlPath}/*.sql" -dir=${modelPath}


  cd ${projectPath} && \
  goctl model mysql datasource --home $GOCTL_TEMPLATE -url="root:xl123456?@tcp(47.105.50.31:3306)/holy-admin" -table="*" -dir=${modelPath}
}

rebuild
