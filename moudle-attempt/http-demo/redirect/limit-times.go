package redirect

import (
	"errors"
	"fmt"
	"http1/basic"
	"net/http"
)

// 限制重定向次数
func RedirectLimitTimes() {
	res, _ := http.NewRequest(http.MethodGet, "http://httpbin.org/redirect/4", nil)
	// 默认的http.DefaultClient len(via) >= 10 10次以上会限制重定向
	// 如果自定义次数需要自己实现一个http.Client
	client := &http.Client{CheckRedirect: func(req *http.Request, via []*http.Request) error {
		if len(via) > 5 {
			return errors.New("redierect too times")
		} else {
			return nil
		}
	}}
	r, err := client.Do(res)
	basic.EexamineErr(err)
	defer r.Body.Close()
	fmt.Println(r.Request.URL)
}
