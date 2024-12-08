package timeout

import (
	"context"
	"io"
	"net"
	"net/http"
	"os"
	"time"
)

// Timeout 演示了如何设置HTTP客户端的超时时间，并处理超时情况
func Timeout() {
	// 创建一个http.Client实例，并设置整个请求的超时时间为5秒
	// 这意味着从发起请求到接收完整响应的整个过程不能超过5秒
	client := &http.Client{

		Timeout: 5 * time.Second, // 设置超时时间
		Transport: &http.Transport{
			// 设计建立连接时间，超时时间被设置为2秒。这意味着如果连接在2秒内没有建立成功，net.DialTimeout将返回一个错误
			DialContext: func(ctx context.Context, network, addr string) (conn net.Conn, e error) {
				return net.DialTimeout(network, addr, 2*time.Second)
			},
			IdleConnTimeout:     2 * time.Second, // 空闲连接的超时时间（与请求超时不同）
			TLSHandshakeTimeout: 2 * time.Second, // TLS握手的超时时间（也与请求超时不同）
		},
	}

	// 使用client发起一个GET请求到httpbin.org/delay/10，该URL会延迟10秒才返回响应
	// 这将触发超时，因为我们的请求超时时间设置为5秒
	res, err := client.Get("https://httpbin.org/delay/10")
	if err != nil {
		// 如果在超时时间内无法建立连接或发送请求，err将不为nil
		// 这里将因为超时错误而抛出异常
		panic(err) // 实际开发中，应该优雅地处理错误，而不是使用panic
	}
	defer res.Body.Close() // 确保在函数返回前关闭响应体，避免资源泄露

	// 注意：由于已经通过client.Get获得了响应，因此不需要再次使用client.Do发送请求

	// 尝试将响应体内容复制到标准输出
	// 但是，由于请求已经超时，res.Body可能已经关闭或处于不可读状态
	// 这里为了示例，我们忽略了io.Copy可能返回的错误
	// 在实际开发中，应该检查并处理这个错误
	_, _ = io.Copy(os.Stdout, res.Body) // 忽略可能的错误

	// 如果请求超过五秒，将会发生以下情况：
	// 1. client.Get调用将返回一个错误，该错误指示请求超时
	// 2. 在上面的代码中，我们通过panic(err)来处理这个错误（但这不是最佳实践）
	// 3. 由于超时，响应体(res.Body)可能已经被关闭或处于不确定状态
	// 4. 尝试从res.Body读取数据可能会失败（尽管在上面的示例中我们忽略了这一点）
	//
	// 在实际开发中，您应该：
	// - 使用if err != nil来检查并处理client.Get返回的错误
	// - 如果发生超时，可以根据需要记录日志、重试请求或返回错误给用户
	// - 始终确保在不再需要时关闭响应体，即使发生错误也是如此
}

// 注意：上面的代码示例主要用于教学目的，展示了如何设置和处理HTTP请求的超时。
// 在实际生产代码中，您应该更加谨慎地处理错误和资源释放。
