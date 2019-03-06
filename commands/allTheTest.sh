#!/usr/bin/env bash

#ldirs=$(find . -name '*_test.go' | xargs dirname | sort | uniq)
tests=$(find . -name '*_test.go' | xargs cat | grep '^func Test' | awk -F '[ (]' '{ print $2 }')

for t in $tests; do
  echo "=== $t"
  #go clean -testcache
  #go build -o go-filecoin main.go

  go test -parallel=1  -run $t
done
