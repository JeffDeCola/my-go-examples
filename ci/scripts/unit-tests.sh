#!/bin/bash
# my-go-examples unit-test.sh

set -e -x

# The code is located in /my-go-examples
echo "pwd is: " $PWD
echo "List whats in the current directory"
ls -lat 

# Setup the gopath based on current directory.
export GOPATH=$PWD

# Now we must move our code from the current directory ./my-go-examples to $GOPATH/src/github.com/JeffDeCola/my-go-examples
mkdir -p src/github.com/JeffDeCola/
cp -R ./my-go-examples src/github.com/JeffDeCola/.

# All set and everything is in the right place for go
echo "Gopath is: " $GOPATH
echo "pwd is: " $PWD
cd src/github.com/JeffDeCola/my-go-examples
ls -lat

# RUN unit_tests and it shows the percentage coverage
# print to stdout and file using tee
go test -cover ./... | tee test_coverage.txt
# add some whitespace to the begining of each line
sed -i -e 's/^/     /' test_coverage.txt

# Move to coverage-results directory.
mv test_coverage.txt $GOPATH/coverage-results/.


