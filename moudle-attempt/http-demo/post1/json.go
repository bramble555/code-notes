package post1

import (
	"bytes"
	"encoding/json"
	"http1/basic"
	"net/http"
)

func PostJson() {
	type data struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	user1 := data{"xionger", 19}
	byt, _ := json.Marshal(user1)
	r, _ := http.Post("http://httpbin.org/post", "application/json", bytes.NewBuffer(byt))
	basic.ReadBody(r)
}
