

- GO111MODULE: GOPATH 模式 或者 模块模式

- GOMODCACHE:  设置 modcache 路径,默认在`$GOPATH/pkg/mod`.

- GOINSECURE: 定义了可以 不安全方式(http进行传输,而不是 不要检查模块校验和数据库验证 这些在GONOSUMDB中)获取模块路径.

- GONOPROXY: 定了某些模块路径 直接从 仓库获取,不从代理拉取.如果没有设置,默认的是 GOPRIVATE

- GONOSUMDB: 定了某些模块路径 不检查模块校验和数据库验证.如果没有设置,默认的是 GOPRIVATE

- GOPATH:  在 GOPATH 模式下,模块缓存在src目录下.

  ​			在模块模式下,模块缓存在 GOPATH 目录下的 pkg/mod 模式下.

  ​			如果没有设置,那就在 用户目录下.

- GOPRIVATE: 用于指定哪些模块路径前缀应被视为私有。

- GOPROXY: 设置代理来下载模块缓冲

  - 特殊值：
    off：禁止从任何来源下载模块。
    direct：直接从版本控制系统下载模块，而不是使用模块代理。

- GOSUMDB: 用于 sum 数据库检验和公钥的使用在 URL 上.它默认会使用 sum.golang.org.

  - 特殊值

    如果将 GOSUMDB 设置为 off 或者使用 -insecure 标志调用 go get，那么校验和数据库将不会被咨询。

- GOVCS:  

  如果没有显式设置 GOVCS:

  - 公共模块：对于公共模块（即那些路径不符合 GOPRIVATE 中定义的模式），go 命令只允许使用 git 和 hg（Mercurial）这两种版本控制系统。
  - 私有模块：对于私有模块（即路径与 GOPRIVATE 中的模式相匹配的模块），go 命令可以使用所有已知的版本控制系统

- GOWORK: 如果被设置成了 off,那就不能用 workspace 模式了

  默认情况下, 搜索 go.work 文件,如果有 就用 workspace 模式 ,否则不用.

  



