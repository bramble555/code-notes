package request1

import (
	"fmt"
	"http1/basic"
	"net/http"
	"net/url"
)

// 如何设置Get请求参数
func ByParameter() {
	// 可以直接把请求参数直接写到URL里面，但是不够灵活
	req, err := http.NewRequest(http.MethodGet, "http://httpbin.org/get", nil)
	basic.EexamineErr(err)
	// 增加请求参数
	parameter := make(url.Values)
	parameter.Add("user", "xionger")
	parameter.Add("age", "18")
	fmt.Println(parameter.Encode()) // age=18&user=xionger
	req.URL.RawQuery = parameter.Encode()
	// 此时URL变成了http://httpbin.org/get?age=18&user=xionger
	r, err := http.DefaultClient.Do(req)
	basic.EexamineErr(err)
	basic.ReadBody(r)
}

// 如何定制请求头
func ByHeader() {
	req, err := http.NewRequest(http.MethodGet, "http://httpbin.org/get", nil)
	basic.EexamineErr(err)
	req.Header.Add("User-Agent", "chorme")
	r, err := http.DefaultClient.Do(req)
	basic.EexamineErr(err)
	basic.ReadBody(r)
}
