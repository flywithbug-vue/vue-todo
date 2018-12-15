#!/usr/bin/env bash

cur_date="`date +%Y-%m-%d-%H:%M:%S`"
dif="$1"
if [ -z "dif" ]; then
    echo "commit msg is empty。
    use cur_date as commit info"
    dif=$cur_date
fi

git add .
git commit -m $dif
git push -u origin master
