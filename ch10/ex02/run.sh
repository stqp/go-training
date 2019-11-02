#! /bin/bash
go build
go run main.go -t zip -f ./a.zip
go run main.go -t tar -f ./a.tar
