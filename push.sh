#!/usr/bin/env bash
cur_date="`date +%Y-%m-%d-%H-%s`"

git add .
git commit -m $cur_date
git push -u origin master
