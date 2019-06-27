#!/bin/bash -e
# run.sh

echo " "
echo "************************************************************************"
echo "******************************************************* run.sh (START) *"
echo " "

echo "Run markdown-delimter-doer with htmltable switch"
go run markdown-delimter-doer.go -htmltable -delimeter \$\$ -i input.md -o output.md
# markdown-delimter-doer -htmltable -delimeter \$\$ -i input.md -o output.md 
echo " "

echo "Remember to install"
echo "    go install markdown-delimter-doer.go"

echo "For more information on markdown-delimter-html-table-generator, read my cheat sheet"
echo "    https://github.com/JeffDeCola/my-go-examples/tree/master/basic-programming/markdown-delimter-doer"
echo " "

echo "********************************************************* run.sh (END) *"
echo "************************************************************************"
echo " "