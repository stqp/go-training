#! /bin/bash

if [ $# -eq 0 ];
then
    echo "no arg. end"
    exit
fi

if [ $1 == "start" ];
then
    go build
    echo "start..."
    TZ=US/Eastern ./ex01 -port 8010 &
    echo $! >> ./pid
    TZ=Asia/Tokyo ./ex01 -port 8020 &
    echo $! >> ./pid
    TZ=Europe/London ./ex01 -port 8030 &
    echo $! >> ./pid
fi

if [ $1 == "stop" ] && [ -e "./pid" ];
then
    for p in `cat ./pid`;
    do
        kill $p
    done
    echo > "./pid"
fi