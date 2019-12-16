#!/usr/bin/env sh

for mainfile in `find $PWD/cmd -name main.go`; do
    go run $mainfile
done