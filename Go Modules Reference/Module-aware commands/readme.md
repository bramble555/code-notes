- GO111MODULE=on,`go` 命令会以模块模式运行
- 如果 `GO111MODULE=auto`，则 `go` 命令会在当前目录或任何父目录中存在 `go.mod` 文件时以模块模式运行。在 Go 1.15 及更早版本中，这是默认行为.

- 如果 `GO111MODULE=off`，则 `go` 命令会忽略 `go.mod` 文件并以 GOPATH 模式运行。

  在 module  模式下下，GOPATH 不再定义构建过程中导入的含义，但它仍然存储下载的依赖项（在 `GOPATH/pkg/mod` 中）和已安装的命令（在 `GOPATH/bin` 中，除非设置了 `GOBIN`）

  从 Go 1.16 开始，无论是否存在 `go.mod` 文件，默认启用模块模式。

  

  在 GOPATH 模式下，`go` 命令忽略模块，它会在 vendor 目录和 GOPATH 中查找依赖项

[Build commands](https://golang.google.cn/ref/mod#build-commands)

所有加载包信息的命令都支持模块模式（module-aware）。这包括：

- `go build`
- `go fix`
- `go generate`
- `go install`
- `go list`
- `go run`
- `go test`
- `go vet`

#### -mod

**`-mod`** 标志控制是否允许自动更新 `go.mod` 文件以及是否使用 vendor 目录。

`-mod=mod`：告诉 `go` 命令忽略 vendor 目录，并自动更新 `go.mod`，例如，当导入的包没有由已知模块提供时。

`-mod=readonly`：告诉 `go` 命令忽略 vendor 目录，并且如果需要更新 `go.mod` 文件时报告错误。

`-mod=vendor`：告诉 `go` 命令使用 vendor 目录。在此模式下，`go` 命令不会使用网络或模块缓存。



如果没有显示 `-mod=` 标志，Go 命令会根据默认行为选择模式。具体行为如下：

- **Go 1.16及更高版本：** 默认启用模块模式（module-aware mode）。如果 `go.mod` 文件存在，Go 命令会自动使用该文件来管理依赖，并且如果缺少模块，Go 会从网络下载它们。如果存在 `vendor` 目录，则会优先使用该目录中的依赖项，就像使用 `-mod=vendor` 一样。
- **Go 1.15及以下版本：** 默认启用 `-mod=auto` 模式，即如果当前目录或任何父目录中存在 `go.mod` 文件，Go 命令会使用模块模式，否则会使用 GOPATH 模式。

总结来说，未显示 `-mod=` 时，Go 会根据版本和是否有 `go.mod` 文件来自动选择合适的模式。如果没有 `go.mod` 文件，Go 会回退到 GOPATH 模式。



`go get` 命令只能在 `-mod=mod` 模式下使用，因为该模式允许修改 `go.mod` 文件，适用于下载、更新和修改依赖项。

#### -GOFLAGS

**`-modcacherw`** 标志指示 `go` 命令在模块缓存中创建具有读写权限的新目录，而不是使它们为只读模式。当始终使用此标志时（通常通过在环境中设置 `GOFLAGS=-modcacherw` 或运行 `go env -w GOFLAGS=-modcacherw`），可以使用诸如 `rm -r` 之类的命令删除模块缓存，而无需首先更改权限。`go clean -modcache` 命令可用于删除模块缓存，无论是否使用了 `-modcacherw`

modcache在 这个路径下`%GOPATH%\pkg\mod`

默认的 只读模式

#### -modfile

**`-modfile` 标志的作用**：

如果你提供了 `-modfile` 标志，比如 `-modfile=custom.mod`，Go 命令会使用 `custom.mod` 文件代替默认的 `go.mod` 文件。

例子

```
go build -modfile=test.mod
```

使用 `test.mod` 作为依赖配置，而不是 `go.mod`。当然也会构建一个 test.sum 文件

[Vendoring](https://golang.google.cn/ref/mod#vendoring)

当你没有网络的时候,这个就派上了用场.

`go mod vendor` 命令会在主模块的根目录中创建一个名为 **vendor** 的目录,如果你手动改了 go.mod, 也需要这个命令来更新 vendor 目录下面的 modules.txt 文件

启用 vendor 或者禁止?

**显式启用**：使用 `-mod=vendor` 标志。

**禁用**：使用 `-mod=readonly` 或 `-mod=mod` 标志。

v1.23 安装的时候 默认是 `-mod=`

根据是否有 vendor 目录决定是否启用 vendor 目录

如果启用了,`go build` 和 `go test` 等构建命令会从 vendor 目录加载包，而不是从网络或本地模块缓存中加载。

回顾一下 `-mod=readonly`构建时如果需要新的依赖，会报错，而不是更新 `go.mod` 和 `go.sum`

那么 v1.23 会更新 go.mod 文件啦



[go get](https://golang.google.cn/ref/mod#go-get)

`go get` 支持以下标志：

- `-d` 标志告诉 `go get` 不 build and install，仅管理 `go.mod` 中的依赖项。在使用 `-d` 时，`go get` 只会更新 `go.mod` 中的依赖项，而不进行 build。

  在 Go 1.17 及之前的版本中，`go get` 默认会在更新模块时构建并安装相关包，但从 Go 1.17 开始，`-d` 将始终启用,`go get` **不再直接用于build 和 install**,`go get` 现在主要用于管理 `go.mod` 中的依赖。如果想 build or install 就用 `go install`命令

- `-u` 标志告诉 `go get` 升级直接或间接导入的包所依赖的模块。每个被选中的模块将升级到最新版本，除非已经要求更高版本（如预发布版本）`-u=patch` 标志会将依赖项升级到最新的补丁版本（类似于 `@patch` 查询）。

- `-t` 标志告诉 `go get` 考虑构建测试所需的模块。当同时使用 `-t` 和 `-u` 时，`go get` 会更新测试依赖项。

Examples:

```
# Upgrade a specific module.最新版本
$ go get golang.org/x/net

# Upgrade modules that provide packages imported by packages in the main module. 所有的!!
$ go get -u ./...

# 升级到指定版本
$ go get golang.org/x/text@v0.3.2

# 移除所有依赖这个模块的依赖项
$ go get golang.org/x/text@none
```

[go install](https://golang.google.cn/ref/mod#go-install)

这个命令是安装程序(build and install),而不是安装依赖项(更新go.mod 文件)



go1.16 以前不支持 版本后缀,只能根据 go.mod  或者 GOPATH 来下载版本

从 Go 1.16 开始，如果参数带有版本后缀（如 `@latest` 或 `@v1.0.0`），`go install` 会在模块感知模式下构建包，**忽略当前目录或任何父目录中的 `go.mod` 文件**。这对于在不影响主模块依赖的情况下安装可执行文件非常有用

`go install` 命令构建并安装命令行中指定路径的包。可执行文件（`main` 包）将安装到 `GOBIN` 环境变量指定的目录，默认情况下为 `$GOPATH/bin`,没有 GOPATH 的话,就下载 在 `$GOROOT/bin`

根据上述所描述,在 go1.17以前版本,go get 和 go install是类似的,只不过 go get 会更新依赖项.

综上所述, go 1.17以后, 必须 go install 来进行 build and install,不能再使用 go get 啦!!

[go list -m](https://golang.google.cn/ref/mod#go-list-m)

`-m` 标志使得 `go list` 列出模块而不是包。

模块和包区分不了?

假设模块是 golang.org/x/net, 包就可能是 golang.org/x/net/http,模块包括了许多包!

如果没有指定任何参数，默认列出主模块(项目自己的模块,而不是项目依赖的模块)

参数

- -f 可以格式化输出

  Example.

  ```
  go list -m -f '{{.Path}} {{.Version}}' golang.org/x/net
  输出
  golang.org/x/net v0.1.0
  
  ```

  可以列举许多参数,具体看官网结构体

  其中 `-retracted` 标志是从 Go 1.16 开始添加的

- -all 可以输出当前目录依赖的所有模块

Example

```
# 列出所有 模块路径, 版本 ,更新时间
go list -m -f 'Module Path: {{.Path}}, Version: {{.Version}}, Last Update: {{.Time}}' all             
```

[go mod download](https://golang.google.cn/ref/mod#go-mod-download)

`go mod download` 命令将指定的模块下载到模块缓存中.

[go mod edit](https://golang.google.cn/ref/mod#go-mod-edit)

`go mod edit` 命令提供了一个命令行接口，用于编辑和格式化 `go.mod` 文件，主要供工具和脚本使用。

目前来看作用不大

[go mod graph](https://golang.google.cn/ref/mod#go-mod-graph)

`go mod graph` 打印图中的每条边，每条边占一行。每行包含两个以空格分隔的字段：一个模块的版本和它的一个依赖。每个模块版本以 `path@version` 的形式标识。主模块没有 `@version` 后缀，因为主模块本身没有版本。

Example

```
github.com/bramble555/blog filippo.io/edwards25519@v1.1.0
```

前者依赖后者.

[go mod init](https://golang.google.cn/ref/mod#go-mod-init)

如果后面 没有 module-path

- 使用 git URL
- 使用 GOPATH 路径

如果存在供应工具的配置文件，`init` 会尝试从中导入模块依赖

最好还是 自定义 module-path .

[go mod tidy](https://golang.google.cn/ref/mod#go-mod-tidy)

`go mod tidy` 增加和删除依赖项,确保 `go.mod` 文件与模块中的源代码相匹配,同时更新 go.sum 文件

参数

- -e 使 `go mod tidy` 尝试继续执行，即使在加载包时遇到错误
- -v  添加和删除的依赖项更加详细 给你打印出来.

[go mod vendor](https://golang.google.cn/ref/mod#go-mod-vendor)

`go mod vendor` 命令用于在主模块的根目录下构建一个名为 `vendor` 的目录,包含依赖项.前面已经讲过 vendor 的作用,这里再次重复一下,当没有网络的时候,使用的依赖项.

参数

- `-e`：在加载包时即使遇到错误，也尝试继续执行（自 Go 1.16 起可用）

- `-v`：添加和删除的依赖项和包更加详细 给你打印出来

- `-o`：（自 Go 1.18 起可用）将 vendored 树输出到指定的目录，而不是默认的 `vendor` 目录。参数可以是绝对路径，也可以是相对于模块根目录的路径

  Example

  ```
  go mod vendor -o my_vendor
  ```

  生成 my_vendor 目录(directory)

  [go mod verify](https://golang.google.cn/ref/mod#go-mod-verify)

  `go mod verify` 主要用于 **验证模块缓存中的文件是否被篡改**

  1. **验证模块是否未被修改**：它会检查模块的 `.zip` 文件和提取的目录，确保自下载以来文件的内容没有发生变化
  2. **校验模块的一致性**：它可以发现模块缓存中是否有文件被篡改或意外修改的情况。



[go mod why](https://golang.google.cn/ref/mod#go-mod-why)

`go mod why` 用于展示从主模块到指定包的导入路径(最短路径)，帮助你理解为什么某个包或模块会被依赖

Example

```
$ go mod why golang.org/x/text/language golang.org/x/text/encoding
# golang.org/x/text/language
rsc.io/quote
rsc.io/sampler
golang.org/x/text/language

# golang.org/x/text/encoding
(main module does not need package golang.org/x/text/encoding)
```

参数 

- -m 显示的是模块,而不是包
- -vendor 忽略主模块以外的 go.mod





[go version ](https://go.dev/ref/mod#go-version-m)

```go
# Print Go version used to build go.
$ go version

我的电脑
go version go1.23.3 windows/386

# 打印 build 后的可执行文件版本
go version D:\language\go\projects\path\bin
D:\language\go\projects\path\bin\air.exe: go1.23.3
D:\language\go\projects\path\bin\gomodifytags.exe: go1.23.3
D:\language\go\projects\path\bin\gopls.exe: go1.23.3
D:\language\go\projects\path\bin\gotests.exe: go1.23.3
D:\language\go\projects\path\bin\migrate.exe: go1.23.3
D:\language\go\projects\path\bin\staticcheck.exe: go1.23.3
```

参数

- -m 会显示模块版本



[go clean (https://go.dev/ref/mod#go-clean-modcache)

清空电脑上的 build 和 测试缓存文件,通常位于 GOPATH/pkg/mod.

如果不确定在哪个路径下,  用这个命令 打印 

```
go env GOMODCACHE
D:\language\go\projects\path\pkg\mod\cache
```

参数

- -modcache  删除 **Go 模块缓存**

  通常这些文件默认 只读 提前设置这些文件可读可写 

  ```
  go env -w GOFLAGS=-modcacherw
  ```

[Version queries](https://go.dev/ref/mod#version-queries)

比如 go get  或者 go install   等等命令 @后面的内容就是版本的内容

有几个特殊点:

1. 默认情况下，`go list -m` 命令 **不会显示** 被撤销的版本，因为这些版本已被标记为不再有效。为了 **显示撤销的版本**，你需要显式地使用 `-retracted` 标志
2. **发布版本** 优先于 **预发布版本**.如果版本 `v1.2.2` 和 `v1.2.3-pre` 都可用，`latest` 查询将选择 `v1.2.2`

[Module commands outside a module](https://go.dev/ref/mod#commands-outside)

没有 go.mod  很难,不能 replace,不能 exclude

[go work init](https://golang.google.cn/ref/mod#go-work-init)

在当前目录下 使用 工作区(go work) 模式,使用这个命令的时候,当前的版本直接写入 在 go.work 里面了

参数

- 路径

  go work init 后面的路径就是 使用 go.mod 模式 的路径

[go work edit](https://golang.google.cn/ref/mod#go-work-edit)

看名字就知道了,不用打开文件,直接编辑 go.work 文件

如果未指定文件，`edit` 会在当前目录及其父目录中查找 `go.work` 文件。

参数

- `-fmt`格式化文件
- `-use=path` 和 `-dropuse=path` 标志用于在 `go.work` 文件的模块目录集合中添加或删除 `use` 指令

- `-go=version` 标志用于设置期望的 Go 语言版本。
- `-replace=old[@v]=new[@v]`   增加 replace 指令
- `-dropreplace=old[@v]`  删除 replace 指令
- `-print` 直接输出 go.work 文件
- `-json` json 形式 输出 go.work 文件

[go work use](https://golang.google.cn/ref/mod#go-work-use)

后面+路径,添加路径到 go.work 文件. 那 上面 edit 命令已经能添加了,为啥 还有 use,这个支持递归添加!!

参数

- -r   搜索你指定的路径及其子目录,如果有 go.mod 文件,添加到 go.work 文件

[go work sync](https://golang.google.cn/ref/mod#go-work-sync)

检查 go.work 文件里面 的 use 路径(他们这些模块)的相同依赖项 版本是否一致.如果不一致,根据 MVS,选择其中最大的版本更新到 每个模块的 go.mod 文件.