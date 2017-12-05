#!/bin/bash
echo "Start build macOS 32 bit"
GOOS=darwin GOARCH=386 go build -o macos/32/main
echo "Completed"
echo "Start build macOS 64 bit"
GOOS=darwin GOARCH=amd64 go build -o macos/64/main
echo  "Completed"
echo "Start build linux 32 bit"
GOOS=linux GOARCH=386 go build -o linux/32/main
echo "Completed"
echo "Start build linux 64 bit"
GOOS=linux GOARCH=amd64 go build -o linux/64/main
echo "Completed"
echo "Start build windows 32 bit"
GOOS=windows GOARCH=386 go build -o windows/32/main
echo "Completed"
echo "Start build windows 64 bit"
GOOS=windows GOARCH=amd64 go build -o windows/64/main
echo "Completed"