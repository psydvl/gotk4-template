#!/usr/bin/env bash
set -uex
if [ $# == 0 ]
then
	array=(master ui)
else
	array=("$@")
fi

for BRANCH in "${array[@]}"; do
	git checkout $BRANCH 
	FILTER_BRANCH_SQUELCH_WARNING=1 git filter-branch --prune-empty -f
	git commit --allow-empty -m "Init with gotk4 minimal template psydvl/gotk4-template"
	git push -u origin $BRANCH --force-with-lease
done

git checkout docs
