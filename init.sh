#!/usr/bin/env bash

if [ $# == 0 ]; then
  echo "Requires a parameter"
fi

if [ $# == 1 ]; then
  echo "[init] clean go mod"
  rm go.mod go.sum
  echo "[init] create main file"
  sed -i 's/PROJECT_NAME/'$1'/g' cmd/*
  sed -i 's/PROJECT_NAME/'$1'/g' config/*
  sed -i 's/PROJECT_NAME/'$1'/g' component/*
  sed -i 's/PROJECT_NAME/'$1'/g' .gitignore
  sed -i 's/PROJECT_NAME/'$1'/g' PROJECT_NAME.go
  sed -i 's/PROJECT_NAME/'$1'/g' infra/error.go
  sed -i 's/PROJECT_NAME/'$1'/g' config_local.yaml
  mv PROJECT_NAME.go $1.go
  go fmt .
  go mod init $1
  go mod tidy
  go build
  rm init.sh
  rm .git -rf
  newapi gen --config api.yaml
  echo "# "$1 > README.md
  echo "[init] All done"
fi

if [ $# -gt 1 ]; then
  echo "Too many parameters"
fi
