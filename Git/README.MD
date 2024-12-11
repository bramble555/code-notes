







切换分支到 dev

```
git check dev
```

如果 当前的本地分支 `dev` 没有设置与远程仓库的跟踪关系

```
git push --set-upstream origin dev
```

`--set-upstream`：用于设置当前本地分支与远程分支的关联。

`origin`：远程仓库的名称（默认情况下为 `origin`）。

`dev`：远程分支的名称（与你的本地分支同名）。



当前分支为 main, 要合并 dev,然后再 推送(push)

记得先拉去 最新的 main 分支

```
git pull origin main

git merge dev
git push origin main

```



