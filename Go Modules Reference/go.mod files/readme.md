[Lexical elements](https://golang.google.cn/ref/mod#go-mod-file-lexical)

keywords

```
`module`,`go`, `require`, `replace`, `exclude`, and `retract
```

[Module paths and versions](https://golang.google.cn/ref/mod#go-mod-file-ident)

- 路径不能以`/`或者`.`开始或结束

[Grammar](https://golang.google.cn/ref/mod#go-mod-file-grammar)

[module directive](https://golang.google.cn/ref/mod#go-mod-file-module)

定义一个路径

Example:

```
module golang.org/x/net

```

**Deprecation**

模块可以通过在包含字符串 Deprecated:（区分大小写）的注释块中标记为已弃用，该字符串位于段落的开头.

Example:

```
// Deprecated: 请使用 example.com/mod/v2 替代。
module example.com/mod
```

不能单独弃用个别次要版本和补丁版本；对于这种情况，使用 retract 可能更为合适。可以用在 v2 及其以上版本

自 Go 1.17 起，go list -m -u 会检查构建列表中所有被标记为弃用的模块的信息。go get 会检查命令行中指定的包构建所需的所有弃用模块。

```
go list -m -u 
github.com/bramble555/blog
```

[go directive](https://golang.google.cn/ref/mod#go-mod-file-go)

表示使用的版本

```
go 1.23.2
```

在 go 1.21之前,表示 建议性的,1.21及其以后表示强制性的要求.

这个作用是选择运行的工具链.

如果没有这个 go directive, 默认 1.16



go 1.21 及其以后 

- 自己项目使用的版本必须大于等于依赖项的版本

- go 命令不再尝试保持与旧版本 Go 的兼容性

啥意思?  你运行项目是1.23,你电脑下载的是 1.21,然后你 go build

```
go build
go: this module requires Go 1.23 or later (module requires Go 1.23)
```

[toolchain directive](https://golang.google.cn/ref/mod#go-mod-file-toolchain)

toolchain 指令用于声明一个建议使用的 go  工具链版本，该工具链将与模块一起使用.相当于 go 指令的直接建议.

不能低于 go 指定的版本

如果没有, 那就和 go  指定的版本一样

同样,你运行项目是1.23,你电脑下载的是 1.21,然后你 go build

```
go build
go: this module requires Go 1.23 or later (module requires Go 1.23)
```

[godebug directive](https://golang.google.cn/ref/mod#go-mod-file-godebug)

godebug 指令用于声明 GODEBUG 版本

如果没有, 那就和 go  指定的版本一样

Example:

```
godebug default=go1.21
godebug (
    panicnil=1
    asynctimerchan=0
)
```

[require directive](https://golang.google.cn/ref/mod#go-mod-file-require)

reuqire 指令用于声明 当前项目依赖的最小版本.是最小版本!!根据 MVS,可能依赖项的版本更大,不可能更小!

go 1.16 及其以下版本 ,间接依赖不会显示声明.可能会发生以下情况,

```
module example.com/main

go 1.16

require (
    github.com/example/foo v1.0.0
    github.com/example/bar v1.2.0 // indirect
)

```

go 1.17 及其以上版本,间接依赖会显示声明,并且支持`模块图修剪`和`延迟模块加载`

[tool directive](https://golang.google.cn/ref/mod#go-mod-file-tool)

同样,也会显示声明工具的使用版本.

工具比如有 格式化代码,go test 覆盖率,性能分析等等

(目前没有使用到工具以后再说)

[exclude directive](https://golang.google.cn/ref/mod#go-mod-file-exclude)

`exclude` 指令用于防止 `go` 命令加载指定版本的模块。主要作用是通过排除某些模块版本来影响依赖解析的行为。

go 1.16 以前,如果某个版本被 `exclude` 指令排除，但仍被引用，`go` 命令会尝试加载更高版本（按版本号排序）来替代它。(不能是伪版本)

这个意思是如果exclude 最新版本,但是仍然被其他模块 require, 那就会error

**Go 1.16 及之后**

- 如果 `exclude` 指令排除了某个版本，而该版本被其他模块的 `require` 指令引用，`go` 命令会忽略对该版本的引用。
- 这可能导致命令（如 `go get` 和 `go mod tidy`）自动添加对更高版本的依赖，并标记为 `// indirect`

什么意思?  我和别人的项目 都依赖了require  example.com/module v1.2.0 ,但是我现在 exclude example.com/module v1.2.0 , go mod tidy 之后, 就变成了  require example.com/module v1.3.0 // indirect

如果有更高的版本会变成这样,否则 error

[replace directive](https://golang.google.cn/ref/mod#go-mod-file-replace)

**replace 指令**用于替换特定版本的模块内容，或者替换所有版本的模块内容。替换的内容可以是另一个模块路径和版本，或者是一个平台特定的文件路径

example:

- 这个指令将替换 `golang.org/x/net` 模块的所有版本，使用 `example.com/fork/net` 的 `v1.4.5` 版本。

```
replace golang.org/x/net => example.com/fork/net v1.4.5

```

- 这个指令将替换 `golang.org/x/net` 的 `v1.2.3` 版本，使用本地路径 `./fork/net` 作为替代，且该路径必须包含一个 `go.mod` 文件。

```
replace golang.org/x/net v1.2.3 => ./fork/net

```

其他类似

[retract directive](https://golang.google.cn/ref/mod#go-mod-file-retract)

**retract 指令**表示某个模块的版本或版本范围不应再被依赖。`retract` 指令用于标记某个版本或版本范围存在问题，或者在该版本发布后发现了严重问题，需要防止用户再依赖该版本

example:

假设模块 `example.com/m` 的作者错误地发布了 `v1.0.0` 版本。为了防止用户继续升级到 `v1.0.0`，作者可以在 `go.mod` 中添加两个撤回指令，并发布 `v1.0.1` 版本，标记为撤回：

```
retract (
    v1.0.0 // Published accidentally.
    v1.0.1 // Contains retractions only.
)
相当于
retract [v0.0.0, v1.0.1] // assuming v1.0.1 contains this 
```

`retract` 指令是从 Go 1.16 开始引入的

[Automatic updates](https://golang.google.cn/ref/mod#go-mod-file-updates)

大多数命令在 `go.mod` 文件缺失信息或与实际情况不符时会报告错误。`go get` 和 `go mod tidy` 命令可以用来修复大多数此类问题。

自动更新 go.mod 更新什么?

- 版本号符合规则
- 遵循排除指令
- 删除冗余或误导的依赖
- 格式化

在 go 1.16以前，`-mod=mod` 标志是默认启用的，因此更新会自动执行。

go 1.16 及其以后,`-mod=readonly` 是默认启用的.只读