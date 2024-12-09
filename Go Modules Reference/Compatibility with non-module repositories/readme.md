你依赖的项目没有 go.mod 文件

1. 从 github 下载

   `go` 命令会在模块缓存中合成一个 `go.mod` 文件，该文件仅包含一个 `module` 指令，不包含任何 `require` 指令。

   但是这个项目可能有间接依赖,就需要需要额外的 `require` 指令（带有 `// indirect` 注释）

2. 从 代理下载

代理会提供一个合成的 `go.mod` 文件

[+incompatible versions](https://golang.google.cn/ref/mod#incompatible-versions)

比如你依赖项没有 go.mod 文件,并且没有遵守 版本规则

会在 依赖项后面 +incompatible

example:

```
require example.com/m v4.1.2+incompatible

```



[Minimal module compatibility](https://golang.google.cn/ref/mod#minimal-module-compatibility)

在 Go 1.11 中引入了“最小模块兼容性”，并且这个功能被回溯到 Go 1.9.7 和 1.10.3。

依赖项在 $GOPATH/src/example.com/repo

如果 v2及其以上,在 $GOPATH/src/example.com/repo/v2

如果依赖项没有遵守版本规则,

路径会被解析为 $modpath/$vn/$dir,是它满足版本规则.