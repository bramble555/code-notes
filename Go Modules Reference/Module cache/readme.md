

module cache 是存储 下载 module 的目录(仅有只读权限).它和 build cache 不同,build cache 包含了已经编译的包和 其他 build 工作.

默认路径 是 $GOPATH/pkg/mod.命令不能删除文件.文件大小没有限制,go 命令不会删除模块,缓存可以被同一台机器上开发的多个 Go 项目共享.

可以使用 `go clean -modcache` 命令来删除缓存。或者，使用 `-modcacherw` 标志时，`go` 命令会以读写权限创建新目录。