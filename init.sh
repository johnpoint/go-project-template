#!/usr/bin/env bash

if [ $# == 0 ]; then
  echo "Requires a parameter"
fi

if [ $# == 1 ]; then
  echo "[init] clean go mod"
  rm go.mod go.sum
  echo "[init] create folder"
  mkdir app app/controller cmd config dao dao/mongoDao model model/mongodb
  echo "[init] create main file"
  sed -i 's/PROJECT_NAME/'$1'/g' app/controller/*
  sed -i 's/PROJECT_NAME/'$1'/g' cmd/*
  sed -i 's/PROJECT_NAME/'$1'/g' config/*
  sed -i 's/PROJECT_NAME/'$1'/g' dao/*
  sed -i 's/PROJECT_NAME/'$1'/g' dao/mongoDao/*
  sed -i 's/PROJECT_NAME/'$1'/g' depend/*
  sed -i 's/PROJECT_NAME/'$1'/g' model/*
  sed -i 's/PROJECT_NAME/'$1'/g' model/mongodb/*
  sed -i 's/PROJECT_NAME/'$1'/g' .gitignore
  sed -i 's/PROJECT_NAME/'$1'/g' PROJECT_NAME.go
  sed -i 's/PROJECT_NAME/'$1'/g' config_local.yaml
  mv PROJECT_NAME.go $1.go
  go fmt $1.go
  go fmt .
  go mod init $1
  go mod tidy
  go build
  rm init.sh
  rm .git -rf
  echo "# "$1 > README.md
  echo "[init] All done"
fi

if [ $# -gt 1 ]; then
  echo "Too many parameters"
fi
