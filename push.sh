#!/usr/bin/env bash
cur_date="`date +%Y-%m-%d-%H:%M:%S`"
dif="$1"
if [ -z "dif" ]; then
    dif=$cur_date
fi

git add .
git commit -m $dif
git push -u origin master
