#!/usr/bin/env bash

function rebuildApi() {
    bash ./admin/shell/rebuild-api.sh
}

function rebuildModel() {
    bash ./admin/shell/rebuild-model.sh
}

function start() {
    bash ./admin/shell/start.sh
}


readonly funcName=$1
#echo "函数名长度:${#funcName}"
if [ -z "${funcName}" ]; then
  echo -e "\033[31m 请输入要执行的函数,例如:start,rebuildApi等\n \033[0m"
else
  "$@"
fi