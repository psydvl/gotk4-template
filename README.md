## Simple mimimal gtk4 + golang template

To use start with:

``` shell
PROJECT="."
#PROJECT="Project_name" # to init in new child directory instead
git clone --branch=master --depth=1 https://github.com/psydvl/gotk4-template $PROJECT
```

### How master branch creating:

``` shell
git checkout changelog
git branch -D master
git switch --orphan master
git commit --allow-empty -m "Init with gotk4 minimal template psydvl/gotk4-template"
git merge --squash --allow-unrelated-histories changelog
git commit --amend --no-edit
git push -u origin master --force-with-lease
```