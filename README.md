## Simple minimal gtk4 + golang template

To use start with:

``` shell
#GOTK_PROJECT="project_name" # set to init in new child directory instead of current
git clone --branch=master --depth=1 https://github.com/psydvl/gotk4-template ${GOTK_PROJECT:-.}
```

## Simple minimal gtk4 + golang template with `.ui` file from [Cambalache](https://flathub.org/apps/details/ar.xjuan.Cambalache)

To use start with:

``` shell
#GOTK_PROJECT="project_name" # set to init in new child directory instead of current
git clone --branch=ui --depth=1 https://github.com/psydvl/gotk4-template ${GOTK_PROJECT:-.}
cd ${GOTK_PROJECT:-.}
git branch -m ui master
```

### How master/ui branches are updating:

``` shell
git checkout docs
# ./ update.sh master ui
./update.sh
```
or
``` shell
BRANCH=master
BRANCH=${BRANCH:-ui} # set to ui if empty, just copy without first line to use
CHANGELOG="changelog-$BRANCH"

git checkout $CHANGELOG
git branch -D $BRANCH
git switch --orphan $BRANCH
git commit --allow-empty -m "Init with gotk4 minimal template psydvl/gotk4-template"
git merge --squash --allow-unrelated-histories $CHANGELOG
git commit --amend --no-edit
git push -u origin $BRANCH --force-with-lease
```