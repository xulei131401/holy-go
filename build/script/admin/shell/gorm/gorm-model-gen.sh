#!/usr/bin/env bash

source ../common.sh

function gen() {
  gormt -g=false
  gormt -H=47.105.50.31 -d=holy-admin -p=xl123456? -u=root --port=3306
}
