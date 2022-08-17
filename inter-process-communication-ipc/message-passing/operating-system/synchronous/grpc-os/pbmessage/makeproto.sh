#!/bin/bash
# makeproto.sh

echo " "
echo "**********************************************************************"
echo "*********************************************** makeproto.sh (START) *"
echo " "

echo "Running protoc"
protoc --go_out=. \
       --go-grpc_out=. \
       --go_opt=paths=source_relative \
       --go-grpc_opt=paths=source_relative \
         pbmessage.proto

echo "*************************************************** makeproto.sh (END) *"
echo "************************************************************************"
echo " "