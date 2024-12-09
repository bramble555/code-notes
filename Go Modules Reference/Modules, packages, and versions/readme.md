modules 可以从仓库下载,也可以从 module 代理 下载.(防止别人把仓库删掉,你的代码不能运行)

module 在 go.mod 文件里面 

[Module paths](https://go.dev/ref/mod#module-path)

路径：可能包含版本（v2及其以上）

[Versions](https://go.dev/ref/mod#versions)

版本：主.次.补丁

- 主版本递增后（每个主版本不兼容），次版本和补丁版本从0开始。例如一个包被移除。
- 次版本递增后（向后兼容更改），补丁版本从0开始。例如一个新的函数被增加。
- 补丁版本递增。例如bug修复或优化。
- 预发布版本。例如v1.2.3-pre 在 v1.2.3 之前。

如果一个版本的主版本号是0，或者它有一个预发布的后缀，则被认为是不稳定的。不稳定的版本不受兼容性要求的限制。

go命令会自动将不遵循此标准的修订版名称转换为规范的版本号。此过程可能会产生一个伪版本。

在主模块之外的go.mod文件里面需要出现规范版本，否则会报错。

[Pseudo-versions](https://go.dev/ref/mod#pseudo-versions)

这里不讨论伪版本

[Major version suffixes](https://go.dev/ref/mod#major-version-suffixes)

如果example.com/mod 是v1.0.0，example.com/mod/v2是v2.0.0

如果旧包和新包有相同的导入路径，则新包必须向后兼容旧包。

模块的新的主版本中的包不向后兼容之前的主版本中的相应包。因此，从v2开始，包需要新的导入路径。主版本后缀在主版本v0或v1时是不允许的。（没必要)



特殊的   gopkg.in/开头的，后面必须有版本（包括v0和v1）例如gopkg.in/yaml.v2

[Resolving a package to a module](https://go.dev/ref/mod#resolve-pkg-mod)

`go`命令正在寻找一个提供包`golang.org/x/net/html`的模块，并且`GOPROXY`被设置为`https://corp.example.com,https://proxy.golang.org`。`go`命令可能会发出以下请求：

- 向 `https://corp.example.com/` （并行）:
- 请求`golang.org/x/net/html`的最新版本
- 请求`golang.org/x/net`的最新版本
- 请求`golang.org/x`的最新版本
- 请求`golang.org`的最新版本

如果都失败了，再向https://proxy.golang.org这个URL请求….

`// indirect`的注释表示该模块是间接依赖.啥意思?

举个例子，假设有三个模块：

1. **主模块 A**（你的项目）

2. **模块 B**，是主模块 A 的直接依赖

3. **模块 C**，是模块 B 的直接依赖

   ```
   require (
       github.com/some/moduleB v1.0.0
       github.com/another/moduleC v2.0.0 // indirect
   )
   
   ```

   