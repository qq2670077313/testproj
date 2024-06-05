#!/usr/bin/env bash

set -e
# echo "" > coverage.txt
cat /dev/null > coverage.txt
cat /dev/null > coverage.out
cat /dev/null > profile.out
START_TIME=`date +%s`

# for d in $(go list ./... | grep -v aaa); do
# for d in $(go list ./...); do
    # go test -short -v -coverprofile=coverage.out -covermode=count "$d"
    go test -short -v -coverprofile=coverage.out -covermode=count "github.com/huof6829/testproj/hello/aaa"
    courtney -l=coverage.out -o=profile.out
    if [ -f profile.out ]; then
        cat profile.out >> coverage.txt
        ## rm profile.out coverage.out
    fi
# done
END_TIME=`date +%s`
EXECUTING_TIME=`expr $END_TIME - $START_TIME`
echo $EXECUTING_TIME
