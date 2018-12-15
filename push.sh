#!/usr/bin/env bash

info() {
     local green="\033[1;32m"
     local normal="\033[0m"
     echo -e "[${green}INFO${normal}] $1"
}

cur_date="`date +%Y-%m-%d-%H:%M:%S`"
dif=$*
if [ -z "dif" ]; then
    echo "commit msg is empty。
    use cur_date as commit info"
    dif=$cur_date
fi

git add .
git commit -m $dif
git push -u origin master
