#!/bin/bash
# my-go-examples makeproto.sh

echo " "
echo "**********************************************************************"
echo "*********************************************** makeproto.sh (START) *"
echo " "

echo "protoc --go_out=. messages.proto"
protoc --go_out=. messages.proto
echo " "

echo "cp up messages.pb.go ../."
cp messages.pb.go ../.
echo ""

echo "*************************************************** makeproto.sh (END) *"
echo "************************************************************************"
echo " "