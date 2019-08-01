#! /bin/bash
go build \
&& ./ex11 -command edit -title "test2" -body "this is body." \
&& ./ex11 -command get -id 1 \
&& ./ex11 -command edit -id 1 -title "changed" \
&& ./ex11 -command close -id 1