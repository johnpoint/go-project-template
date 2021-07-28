#!/usr/bin/env bash

if [ $# == 0 ]; then
  echo "Requires a parameter"
fi

if [ $# == 1 ]; then
  echo "[init] clean go mod"
  rm go.mod go.sum
  echo "[init] create folder"
  mkdir app app/controller cmd config dao dao/mongoDao initHelper initHelper/depend model model/mongodb
  echo "[init] create main file"
  sed -i 's/{{project_name}}/'$1'/g' app/controller/*
  sed -i 's/{{project_name}}/'$1'/g' cmd/*
  sed -i 's/{{project_name}}/'$1'/g' config/*
  sed -i 's/{{project_name}}/'$1'/g' dao/*
  sed -i 's/{{project_name}}/'$1'/g' dao/mongoDao/*
  sed -i 's/{{project_name}}/'$1'/g' initHelper/*
  sed -i 's/{{project_name}}/'$1'/g' initHelper/depend/*
  sed -i 's/{{project_name}}/'$1'/g' model/*
  sed -i 's/{{project_name}}/'$1'/g' model/mongodb/*
  echo 'package main
import "'$1'/cmd"
func main(){cmd.Execute()}
' > $1.go
  go fmt $1.go
  go fmt .
  go mod init $1
  go mod tidy
  go build
  echo "[init] All done"
fi

if [ $# -gt 1 ]; then
  echo "Too many parameters"
fi
