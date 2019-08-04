#!/bin/bash -e
# run.sh

echo " "
echo "************************************************************************"
echo "******************************************************* run.sh (START) *"
echo " "

echo "Run markdown-create-table-of-contents"
go run markdown-create-table-of-contents.go -i input.md -h3
echo " "

echo "Remember to install"
echo "    go install markdown-create-table-of-contents.go"

echo "For more information on markdown-delimter-html-table-generator, read my cheat sheet"
echo "    https://github.com/JeffDeCola/my-go-examples/tree/master/useful-tools-i-use/markdown-create-table-of-contents"
echo " "

echo "********************************************************* run.sh (END) *"
echo "************************************************************************"
echo " "