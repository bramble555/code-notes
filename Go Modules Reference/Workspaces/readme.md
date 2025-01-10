大部分命令都是在 一个 工作区执行

`go.work` 文件）在 Go 1.18 及以后版本中被引入，主要用于管理多个 Go 模块并进行模块之间的协调与依赖管理.

可以使用 `go env GOWORK` 命令来查看 Go 命令正在使用哪个 `go.work` 文件。如果 Go 命令不在工作空间模式下，`go env GOWORK` 将返回空值

如果 `GOWORK` 为空或未设置，命令会在当前工作目录及其父目录中查找 `go.work` 文件。如果找到该文件，命令将在它定义的工作空间中执行；

否则，工作空间将只包含包含当前工作目录的模块。

```
/project
    /moduleA
        go.mod
        main.go
    /moduleB
        go.mod
        helper.go
    go.work
	
这是我的项目
go.work 文件 内容如下
__________________________________
go 1.18

use (
    ./moduleA
    ./moduleB
)


_________________________________
// moduleA/main.go
package main

import (
    "fmt"
    "moduleB"  // 直接引用 moduleB 模块
)

func main() {
    fmt.Println(moduleB.HelperFunction())  // 使用 moduleB 中的函数
}


```

这样可以 多模块协作,提高开发效率,避免冗余依赖

notes: 整个工作空间只能使用一个 Go 版本

[go.work files](https://golang.google.cn/ref/mod#go-work-file)

Go 命令提供了几个子命令来操作 `go.work` 文件。`go work init` 用于创建新的 `go.work` 文件，`go work use` 用于向 `go.work` 文件中添加模块目录，`go work edit` 用于进行低级编辑。Go 程序还可以通过 `golang.org/x/mod/modfile` 包来编程修改 `go.work` 文件。

Go 命令将维护一个 `go.work.sum` 文件，跟踪工作空间中那些不在集体工作空间模块的 `go.sum` 文件中的哈希值。

不建议 git 提交 go.work 文件,可能对其他开发者有影响,啥时候提交呢?彼此模块之间确实要依赖.

[Lexical elements](https://golang.google.cn/ref/mod#go-work-file-lexical)

和 go.mod 文件类似,不再过多描述

只是 keywords 为 `use`,后面 是 路径

[go directive](https://golang.google.cn/ref/mod#go-work-file-go)

```
GoWork = { Directive } .
Directive = GoDirective |
            ToolchainDirective |
            UseDirective |
            ReplaceDirective .
```

当然也有一个 go 指令,后面+版本号

下面详细描述

[toolchain directive](https://golang.google.cn/ref/mod#go-work-file-toolchain)

`toolchain` 指令用于指定建议的 Go 工具链版本。如果当前使用的默认工具链版本较旧，则该指令会生效.

和 go.mod 文件类似,没有指定  那就和 go 指令版本一样

[godebug directive](https://golang.google.cn/ref/mod#go-work-file-godebug)

其语法和效果与 `go.mod` 文件中的 `godebug` 指令相同。

当使用 go.work 时，go.mod 文件中的 godebug 指令会被忽略，只有在 go.work 文件中的 godebug 指令才会生效。

[use directive](https://golang.google.cn/ref/mod#go-work-file-use)

它的参数是一个相对路径，指向包含模块 `go.mod` 文件的目录。

[replace directive](https://golang.google.cn/ref/mod#go-work-file-replace)

是的，`replace` 指令替换的是 **模块**

在 `go.work` 文件中使用了 `replace` 指令，那么它会覆盖工作空间内其他模块（即在工作空间中的各个 `go.mod` 文件中）对相同模块或相同版本的 `replace` 指令

简单来说, go.work 的 replace  指令 优先级 大于  go.mod 的 replace指令(相同模块下)