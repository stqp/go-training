#! /bin/sh
go run ex12/main.go web & 
ppid=$(echo $!)
sleep 2
curl -s -o ex12/out.gif http://localhost:8000/?cycle=20&nframe=100
sleep 2
pkill -P ${ppid}