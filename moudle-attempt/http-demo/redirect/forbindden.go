package redirect

import (
	"fmt"
	"http1/basic"
	"net/http"
)

func RedirectForbindden() {
	// 禁止重定向
	// 登录请求，防止重定向到首页
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	res, _ := http.NewRequest(http.MethodGet, "http://httpbin.org/redirect/4", nil)
	r, err := client.Do(res)
	basic.EexamineErr(err)
	defer r.Body.Close()
	fmt.Println(r.Request.URL)
}
