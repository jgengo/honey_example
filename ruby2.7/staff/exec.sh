#!/bin/sh

if [ "$#" -lt 2 ]; then
    echo "usage: <ex number> <files ...>"
    echo "example: ex00 main.rb"
    exit 1
fi

ruby user_dir/$1/$2 > user_output.txt

diff answers/$1/output.txt user_output.txt | cat -e
if [ $? -eq 0 ]
then
    echo "OK"
else
    echo "KO"
fi
