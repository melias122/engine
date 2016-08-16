@echo off

GOOS=windows GOARCH=386 go build -a -ldflags="-H windowsgui" -o 386.exe
GOOS=windows GOARCH=amd64 go build -a -ldflags="-H windowsgui" -o amd64.exe

