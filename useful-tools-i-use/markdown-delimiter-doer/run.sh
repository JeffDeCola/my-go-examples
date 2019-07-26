#!/bin/bash -e
# run.sh

echo " "
echo "************************************************************************"
echo "******************************************************* run.sh (START) *"
echo " "

echo "Run markdown-delimiter-doer with htmltable switch"
go run markdown-delimiter-doer.go -htmltable -delimiter \$\$ -i input.md -o output.md
# markdown-delimiter-doer -htmltable -delimiter \$\$ -i input.md -o output.md 
echo " "

echo "Remember to install"
echo "    go install markdown-delimiter-doer.go"

echo "For more information on markdown-delimter-html-table-generator, read my cheat sheet"
echo "    https://github.com/JeffDeCola/my-go-examples/tree/master/useful-tools-i-use/markdown-delimiter-doer"
echo " "

echo "********************************************************* run.sh (END) *"
echo "************************************************************************"
echo " "