package response

import (
	"bufio"
	"fmt"
	"http1/basic"
	"io"
	"net/http"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/transform"
)

// 响应信息
func Info() {
	r, err := http.Get("https://baidu.com")
	basic.EexamineErr(err)
	defer r.Body.Close()
	fmt.Println("响应状态码为", r.StatusCode)
	fmt.Println("响应信息为", r.Status)
	fmt.Println("响应头Content-Type信息为", r.Header.Get("Content-Type")) // 忽略大小写,如果使用r.Header["....."] 不忽略大小写
	// 如何获取编码信息？
	// content-tpye 可能会提供编码
	// html head meta 可能获取编码
	// 以上都不行，可以通过网页头部信息猜网页编码信息

	// b, err := io.ReadAll(r.Body)
	// basic.EexamineErr(err)
	// 如果不想读取body
	bufferReader := bufio.NewReader(r.Body)
	bytes, _ := bufferReader.Peek(1024)
	e, name, _ := charset.DetermineEncoding(bytes, r.Header.Get("Content-Type"))
	fmt.Println("编码为", e, "编码名字为", name)
	// 如果不是utf-8，转化为utf-8编码，这个转换函数参数为bufferReader和转化器（解码）
	bodyReader := transform.NewReader(bufferReader, e.NewDecoder())
	content, _ := io.ReadAll(bodyReader)
	fmt.Println(string(content))

}
