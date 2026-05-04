#!/bin/bash
rm -rf build

GOOS=windows GOARCH=amd64 go build -o build/local-finder-app.exe
GOOS=linux GOARCH=amd64 go build -o build/local-finder-app-linux
GOOS=darwin GOARCH=amd64 go build -o build/local-finder-app-mac
GOOS=darwin GOARCH=arm64 go build -o build/local-finder-app-mac-arm