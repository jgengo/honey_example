#!/bin/sh

if [ "$#" -lt 2 ]; then
    echo "no"
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
