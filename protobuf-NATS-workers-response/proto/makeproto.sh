#!/bin/bash
# my-go-examples makeproto.sh

protoc --lint_out=. *.proto
protoc --go_out=. messages.proto
cp messages.pb.go ../server/.
cp messages.pb.go ../client/.