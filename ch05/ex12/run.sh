#! /bin/bash
url="https://web.ics.purdue.edu/~gchopra/class/public/pages/webdesign/05_simple.html"
go run main.go $url > actual
dif=`diff expected actual`
if [ ${#dif} -gt 0 ]
then
    echo "Failed."
    echo "Diff:"$dif
    exit 1
fi
echo "OK"
