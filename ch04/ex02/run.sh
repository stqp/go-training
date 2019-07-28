#! /bin/bash
echo "test"  | go run main.go -t 256
echo "test"  | go run main.go -t 384
echo "test"  | go run main.go -t 512