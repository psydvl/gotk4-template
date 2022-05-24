#!/usr/bin/env bash
set -uex
if [ $# == 0 ]
then
	array=(master ui)
else
	array=("$@")
fi

for BRANCH in "${array[@]}"; do
	CHANGELOG="changelog-$BRANCH"
	git checkout $CHANGELOG
	git branch -D $BRANCH
	git switch --orphan $BRANCH
	git commit --allow-empty -m "Init with gotk4 minimal template psydvl/gotk4-template"
	git merge --squash --allow-unrelated-histories $CHANGELOG
	git commit --amend --no-edit
	git push -u origin $BRANCH --force-with-lease || true # in case no network avaliable
done

git checkout docs
