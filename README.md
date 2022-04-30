## Simple minimal gtk4 + golang template

To use start with:

``` shell
#GOTK_PROJECT="Project_name" # set to init in new child directory instead of current
git clone --branch=master --depth=1 https://github.com/psydvl/gotk4-template ${GOTK_PROJECT:-.}
```

## Simple minimal gtk4 + golang template with `.ui` file from [Cambalache](https://flathub.org/apps/details/ar.xjuan.Cambalache)

To use start with:

``` shell
#GOTK_PROJECT="Project_name" # set to init in new child directory instead of current
git clone --branch=ui --depth=1 https://github.com/psydvl/gotk4-template ${GOTK_PROJECT:-.}
```

### Preview other branches

| [master](../../tree/master) | [ui](../../tree/master) |
| --- | --- |

### How master/ui branch creating:

We let last commit to be empty then command `git clone --depth 1` can take only it with right message locally, coexist with editing history in server

``` shell
BRANCH=master
BRANCH=${BRANCH:-ui} # set to ui if empty, just copy without first line to use
git checkout $BRANCH 
FILTER_BRANCH_SQUELCH_WARNING=1 git filter-branch --prune-empty
git commit --allow-empty -m "Init with gotk4 minimal template psydvl/gotk4-template"
git push -u origin $BRANCH --force-with-lease
```