#!/usr/bin/env bash
cur_date="`date +%Y-%m-%d-%H-%M-%S`"

git add .
git commit -m $cur_date
git push -u origin master
