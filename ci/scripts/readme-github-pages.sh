#!/bin/sh
# my-go-examples readme-github-pages.sh

echo " "

if [ "$1" = "-debug" ]
then
    echo "readme-github-pages.sh -debug (START)"
    # set -e causes the shell to exit if any subcommand or pipeline returns a non-zero status. Needed for concourse.
    # set -x enables a mode of the shell where all executed commands are printed to the terminal.
    set -e -x
    echo " "
else
    echo "readme-github-pages.sh (START)"
    # set -e causes the shell to exit if any subcommand or pipeline returns a non-zero status.  Needed for concourse.
    set -e
    echo " "
fi

echo "The goal is to git clone /my-go-examples to /my-go-examples-updated"
echo "Then script will edit the /docs/_includes/README.md for GITHUB WEBPAGES"
echo "Finally push the changes in /docs/_includes/README.md to github"
echo " "

echo "At start, you should be in a /tmp/build/xxxxx directory with two folders:"
echo "   /my-go-examples"
echo "   /my-go-examples-updated (created in task-build-push.yml task file)"
echo " "

echo "pwd is: $PWD"
echo " "

echo "List whats in the current directory"
ls -la
echo " "

echo "git clone my-go-examples to my-go-examples-updated"
git clone my-go-examples my-go-examples-updated
echo " "

echo "cd my-go-examples-updated"
cd my-go-examples-updated
echo " "

echo "List whats in the current directory"
ls -la
echo " "

echo "FOR GITHUB WEBPAGES"
echo "THE GOAL IS TO COPY README.md to /docs/_includes/README.md"
echo " "

echo "Remove everything before the second heading in README.md.  Place in temp-README.md"
sed '0,/GitHub Webpage/d' README.md > temp-README.md
# Change the first heading ## to #
sed -i '0,/##/{s/##/#/}' temp-README.md
# update the image links (remove docs/)
sed -i 's#IMAGE](docs/#IMAGE](#g' temp-README.md
echo " "

commit="yes"

echo "Does docs/_includes/README.md exist?"
if test -f docs/_includes/README.md
then
    echo "    Yes, it exists."
    # CHECK IF THERE IS A DIFF
    if (cmp -s temp-README.md docs/_includes/README.md)
    then
        commit="no"
        echo "    No changes are needed, Do not need to git commit and push"
    else
        echo "    Updates are needed"
    fi
    echo " "
else
    echo "    No, it does not exist"
    echo "    Creating the _includes directory"
    mkdir docs/_includes
    echo " "
fi

if [ "$commit" = "yes" ]
then
    echo "cp updated temp-README.md to docs/_includes/README.md"
    cp temp-README.md docs/_includes/README.md
    echo " "

    echo "update some global git variables"
    git config --global user.email "jeff@keeperlabs.com"
    git config --global user.name "Jeff DeCola (Concourse)"
    echo " "
    git config --list
    echo " "

    echo "ONLY git add and commit what is needed to protect from unforseen issues"
    echo "git add"
    git add docs/_includes/README.md
    echo " "

    echo "git commit"
    git commit -m "Update docs/_includes/README.md for GitHub WebPage"
    echo " "

    echo "git status"
    git status
    echo " "
    
    echo "git push  - not needed in concourse since its done in pipeline"
    echo " "
fi

echo "remove temp-README.md"
rm temp-README.md
echo " "

echo "readme-github-pages.sh (END)"
echo " "
