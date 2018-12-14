#!/usr/bin/env bash
cur_date="`date +%Y-%m-%d-%H-%m-%s`"

git add .
git commit -m $cur_date
git push -u origin master
