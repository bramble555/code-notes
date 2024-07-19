
package main

import (
	"context"   // 引入context包，用于处理超时和取消信号
	"log"       // 引入log包，用于记录日志
	"net/http"  // 引入net/http包，用于HTTP服务器
	"os"        // 引入os包，用于操作系统交互，如信号处理
	"os/signal" // 引入os/signal包，用于处理信号
	"time"      // 引入time包，用于处理时间

	"github.com/gin-gonic/gin" // 引入Gin框架，用于构建Web应用
)

func main() {
	// 创建Gin路由引擎
	router := gin.Default()
	// 定义一个GET请求路由，当访问根URL时触发
	router.GET("/", func(c *gin.Context) {
		// 模拟长时间处理请求
		time.Sleep(5 * time.Second)
		// 响应HTTP状态码200和欢迎信息
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	// 创建HTTP服务器实例
	srv := &http.Server{
		Addr:    ":8080", // 监听地址和端口
		Handler: router,  // 设置请求处理程序
	}

	// 在新的goroutine中启动HTTP服务器
	go func() {
		// 监听并服务请求
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			// 如果出现错误且不是服务器已关闭的错误，则记录并退出
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 创建一个通道来接收中断信号
	quit := make(chan os.Signal, 1)
	// 通知这个通道当接收到中断信号（如Ctrl+C）
	signal.Notify(quit, os.Interrupt)
	// 阻塞等待中断信号
	<-quit
	log.Println("Shutdown Server ...")

	// 创建一个带有超时时间的context，用于优雅地关闭服务器
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// 在函数结束时取消context，释放资源
	defer cancel()
	// 尝试优雅地关闭服务器，等待当前连接完成或超时
	if err := srv.Shutdown(ctx); err != nil {
		// 如果关闭服务器时出错，则记录并退出
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
