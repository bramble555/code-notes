

支持 Git, Subversion, Mercurial, Bazaar, and Fossil

假如 使用 Git 命令 需要添加到 环境变量里面



[Finding a repository for a module path](https://golang.google.cn/ref/mod#vcs-find)

只会下载特定版本,而不会下载 全部版本.



[Mapping versions to commits](https://golang.google.cn/ref/mod#vcs-version)

每个版本都应该在 GitHub 中 有对应的 tag

根目录下的模块就不用说了

如果模块是在仓库的某个子目录中定义的，也就是说模块路径的子目录部分不是空的，那么每个标签名必须以模块的子目录名作为前缀，并跟一个斜杠

```
例如，模块golang.org/x/tools/gopls是在仓库根路径为golang.org/x/tools下的gopls子目录中定义的。因此，这个模块的v0.4.0版本在仓库中的标签就必须是gopls/v0.4.0。
```

如果不存在go.mod文件，并且模块位于仓库的根目录中，那么主要版本号为v2或更高的标签可以属于一个没有主要版本后缀的模块。会在版本描述中加上+incompatible，但在Git标签本身中不会出现这个后缀。

tag 创建之后不应该删除或者修改,即使别人真的删库跑路了,模块代理里面有.



[Mapping pseudo-versions to commits](https://golang.google.cn/ref/mod#vcs-pseudo)

伪版本号 将会把最后12个字符作为 tag



[Mapping branches and commits to versions](https://golang.google.cn/ref/mod#vcs-branch)



go get 的时候可能会下载特定的 branch

如果有特定的主版本号,会选择 对应的主版本号 和最高次版本

如果没有特定的主版本号,会选择最高的版本

如果提交的时候,部分版本没有 标准的版本号,

比如

```
提交A：v1.0.0
提交B：从A继承而来，无标签
提交C：从B继承而来，无标签
提交D：从C继承而来，带有标签v1.2.0
提交E：从D继承而来，无标签
提交F：从E继承而来，无标签
```

下载 example.com/mod/v1 的时候,会下载 提交D  ,还会加上一些额外的信息（如时间戳和提交哈希值的前缀），生成一个伪版本

[Module directories within a repository](https://golang.google.cn/ref/mod#vcs-dir)

仓库根路径：example.com/myrepo
模块路径：example.com/myrepo

大多数情况是这样的,仓库根路径和模块路径要相同



有时候模块定义在仓库的子目录中。这通常是为了处理大型仓库，这些仓库中有多个组件需要独立发布和版本化。这样的模块预期会位于与模块路径中仓库根路径之后的部分相匹配的子目录中。例如，假设模块example.com/monorepo/foo/bar位于根路径为example.com/monorepo的仓库中。那么它的go.mod文件必须位于foo/bar子目录中。

例如，假设上述模块的新版本路径为example.com/monorepo/foo/bar/v2。那么它的go.mod文件可能位于foo

/bar或foo/bar/v2目录中。





模块的.zip文件不包括vendor目录的内容或任何嵌套模块（包含go.mod文件的子目录）。这意味着模块必须确保不引用其目录外或其它模块中的文件。



如果仓库中有一个testdata目录包含了大文件，模块作者可以在testdata中添加一个空的go.mod文件，这样用户就不必下载那些文件。



[Special case for LICENSE files](https://golang.google.cn/ref/mod#vcs-license)

允许同一个LICENSE文件适用于仓库内的所有模块(不能含有扩展名)

也就是说,子目录下的 go.mod 文件 如果没有 LICENSE文件,下载的时候,也会自动创建一个 LICENSE文件



[Controlling version control tools with GOVCS](https://golang.google.cn/ref/mod#vcs-govcs)

代码可以从任何服务器导入,这是`去中心化`的

为了安全,默认情况下，Go命令只会使用Git和Mercurial (hg) 从公共服务器下载代码。公共模块还是先从代理处下载

当下载私有模块的时候,将直接从仓库(服务器)下载,不能从代理下载.



#### GOVCS变量的格式和规则

`GOVCS`变量由一系列逗号分隔的规则组成，每个规则的形式是`pattern:vcslist`。其中：

- **pattern**：是一个glob模式，用于匹配模块路径或导入路径的一个或多个前导元素。

- vcslist

  ：是由竖线（|）分隔的版本控制命令列表，可以是：

  - 具体的版本控制系统名称（如`git`、`hg`等）
  - `all`：表示允许所有已知的版本控制系统
  - `off`：表示不允许任何版本控制系统

例如

```
GOVCS=github.com:git,evil.com:off,*:git|hg
```

对于所有以github.com/开头的模块路径，只允许使用Git作为版本控制系统。

对于所有以evil.com/开头的模块路径，禁止使用任何版本控制系统。但是，如果源服务器(原来的项目)使用mod方案（即通过GOPROXY协议提供模块），这些模块仍然可以被下载。

对于所有其他模块路径，只允许使用Git或Mercurial作为版本控制系统



可以使用 `go env -w GOVCS="github.com:git,evil.com:off,*:git|hg"  `来进行设置

`GOVCS` 在 GO 1.16被引入







