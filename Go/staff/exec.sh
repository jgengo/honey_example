#!/bin/sh

if [ "$#" -lt 2 ]; then
    echo "usage: <ex number> <files ...>"
    echo "example: ex00 main.c"
    echo "example: ex00 *.c"
    exit 1
fi
go run user_dir/$1/$2 > user_output.txt

diff answers/$1/output.txt user_output.txt &> /dev/null

if [ $? -eq 0 ]
then
    echo "OK"
else
    echo "KO"
fi
