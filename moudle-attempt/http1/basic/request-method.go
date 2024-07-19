package basic

import (
	"fmt"
	"io"
	"net/http"
)

func Get() {
	r, err := http.Get("http://httpbin.org/get")
	if err != nil {
		panic(err)
	}
	// 记得关闭body
	defer r.Body.Close()
	content, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}
func Post() {
	// 比get多了俩个参数
	r, err := http.Post("http://httpbin.org/post", "", nil)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	content, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}

// http没有Put方法，需要自己去写
func Put() {
	req, err := http.NewRequest(http.MethodPut, "http://httpbin.org/put", nil)
	if err != nil {
		panic(err)
	}
	r, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	content, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}

// Delete和put一模一样
func Delete() {

}

// head
// options 同理
func EexamineErr(err error) {
	if err != nil {
		panic(err)
	}
}
func ReadBody(r *http.Response) {
	defer r.Body.Close()
	content, err := io.ReadAll(r.Body)
	EexamineErr(err)
	fmt.Println(string(content))
}
