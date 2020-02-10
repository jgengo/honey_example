#!/bin/bash

if [ "$#" -lt 2 ]; then
    echo "usage: <ex number> <files ...>"
    echo "example: ex00 main.c"
    echo "example: ex00 *.c"
    exit 1
fi
gcc -o output user_dir/$1/${@:2}
./output > user_output.txt

diff answers/$1/output.txt user_output.txt &> /dev/null

if [ $? -eq 0 ]
then
    echo "OK"
else
    echo "KO"
fi
