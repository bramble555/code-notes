

[GOPROXY protocol](https://go.dev/ref/mod#goproxy-protocol)

代理是GET 请求

规则:

1. 支持重定向
2. 4** 和5**是错误

多个代理之间用 `,`或者`|`

**逗号 `,`**：仅在当前代理返回 **`404 (Not Found)` 或 `410 (Gone)`** 时，才会回退到下一个代理。

**竖线 `|`**：只要当前代理出现**任何错误**（包括非 HTTP 错误，比如超时），就会回退到下一个代理。



如果要下载的版本没有,会下载最新的版本.(语义最高的版本)

#### Go 1.12 的选择顺序：

1. 正式版本（如 `v1.3.0`）。
2. 预发布版本（如 `v1.3.0-beta.1`）。
3. 伪版本（如 `v1.3.1-20240101010101-abcdefabcdef`）。

#### Go 1.13 的选择顺序：

1. 语义版本号最高的版本（包括伪版本）。
   - `v1.3.1-20240101010101-abcdefabcdef` 会优先于 `v1.3.0`。



[Communicating with proxies](https://go.dev/ref/mod#communicating-with-proxies)

**环境变量的作用**：

- `GOPRIVATE` 和 `GONOSUMDB`：用于禁用对特定模块的校验和数据库请求。
- `GOSUMDB`：设置为 `off` 时，会完全禁用对校验和数据库的请求。

[Serving modules directly from a proxy](https://go.dev/ref/mod#serving-from-proxy)

大型企业可以在内部搭建一个 `GOPROXY` 代理，用于提供企业专有的 Go 模块，而无需公开其 Git 仓库或切换到标准工具链.



用户运行：

```
go get example.com/gopher@v1.0.0

GET https://modproxy.example.com/example.com/gopher/@v/v1.0.0.info
GET https://modproxy.example.com/example.com/gopher/@v/v1.0.0.mod
GET https://modproxy.example.com/example.com/gopher/@v/v1.0.0.zip
```

- `go` 命令解析 `<meta>` 标签，知道模块由 `https://modproxy.example.com` 提供。
- 依次发送 `.info`、`.mod`、`.zip` 请求获取模块的相关数据和源码。

