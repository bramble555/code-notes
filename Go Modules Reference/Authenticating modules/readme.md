

`go` 命令会将每个哈希值与主模块的 `go.sum` 文件中的对应行进行比较。如果哈希值与 `go.sum` 中的哈希值不同，`go` 命令会报告安全错误，并删除下载的文件，而不会将其添加到模块缓存中。

哈希值验证通过后，`go` 命令会将哈希值添加到 `go.sum` 文件中，并将下载的文件添加到模块缓存中。如果模块是私有的（通过 `GOPRIVATE` 或 `GONOSUMDB` 环境变量匹配），或者校验和数据库被禁用（通过设置 `GOSUMDB=off`），`go` 命令会接受该哈希值并将文件添加到模块缓存中，而不进行验证。



模块缓存通常会被系统上的所有 Go 项目共享，每个模块可能有自己的 `go.sum` 文件，其中包含不同的哈希值。为了避免信任其他模块，`go` 命令在访问模块缓存中的文件时，始终使用主模块的 `go.sum` 文件来验证哈希值。由于计算 `.zip` 文件的哈希值开销较大，`go` 命令会检查与 `.zip` 文件一起存储的预计算哈希值，而不是重新计算哈希。可以使用 `go mod verify` 命令检查 `.zip` 文件和解压目录自添加到模块缓存以来是否已被修改



[go.sum files](https://golang.google.cn/ref/mod#go-sum-files)

go 命令会对下载的 .mod 文件或整个 .zip 文件内容计算其哈希值，然后与 mod cache 里面的 哈希进行对比，如果被篡改过，为了保护项目，就不导入了.

go.sum 文件 包含 路径，版本，哈希值

go mod tidy 也会更新 go.sum 文件

[Checksum database](https://golang.google.cn/ref/mod#checksum-database)

GOSUMDB 默认为由 Google 运行的 Go 校验和数据库 sum.golang.org。

如果 GOSUMDB 设置为 off，或者 go get 使用 -insecure 标志调用，则不会咨询校验和数据库，并接受所有未识别的模块，代价是放弃所有模块经过验证的可重复下载的安全保证。绕过特定模块的校验和数据库的更好方法是使用 GOPRIVATE 或 GONOSUMDB 环境变量。Windows 设置环境变量较麻烦，Linux 使用 export 命令  比如 export GOPRIVATE="mycompany.com/*"







