#! /bin/bash

pkgs=`go list ...`

for pkg in ${pkgs}:
do 
    #echo $pkg
    #echo "go list $pkg"
    $(go list -f '{{join .Deps "\n"}}' $pkg)
done
