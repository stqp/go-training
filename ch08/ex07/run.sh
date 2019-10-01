#! /bin/bash
go build

# run with crawler mode
./ex07 -mode=1 https://golang.org &
pid=$!

# wait a little bit for proceed crawling.
sleep 2
kill $pid


# run with file server mode
./ex07 -mode=2 &
pid=$!

# test response
sleep 3
curl http://localhost:5000/

kill $pid
