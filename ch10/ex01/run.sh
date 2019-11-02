#! /bin/bash
go build

in="img.jpg"

cat $in | ./ex01 -t jpg > out.jpg
cat $in | ./ex01 -t png > out.png
cat $in | ./ex01 -t gif > out.gif

