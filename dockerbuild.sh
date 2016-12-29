#!/bin/bash

export GOOS=linux
export GOARCH=amd64

revel build mysite build

docker build -f Dockerfile -t cpurta/mysite:latest .
