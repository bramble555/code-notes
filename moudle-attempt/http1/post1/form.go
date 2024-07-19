package post1

import (
	"http1/basic"
	"net/http"
	"net/url"
	"strings"
)

func PostForm() {
	// form data（表单数据）形式类似于 query string，name=xionger&age=18
	data := make(url.Values)
	data.Add("name", "xionger")
	data.Add("age", "18")
	dataString := data.Encode()
	r, err := http.Post("http://httpbin.org/post", "application/x-www-form-urlencoded", strings.NewReader(dataString))
	basic.EexamineErr(err)
	basic.ReadBody(r)
}
