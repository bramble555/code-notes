package proxy1

import (
	"io"
	"net/http"
	"net/url"
	"os"
)

func Proxy() {
	// 没设置代理，没办法进行测试
	proxyUrl, _ := url.Parse("http://127.0.0.1:8087")
	t := &http.Transport{
		Proxy: http.ProxyURL(proxyUrl),
	}
	// 一般主要两种，http 代理 和 shadowsocksR(SSR) 的代码, socks5
	client := http.Client{Transport: t}
	r, err := client.Get("https//goggle.com")
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	_, _ = io.Copy(os.Stdout, r.Body)
	// 有的可能会session进行包装
}
