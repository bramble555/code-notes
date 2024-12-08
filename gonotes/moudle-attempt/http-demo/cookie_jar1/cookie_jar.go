package cookiejar1

import (
	"io"
	"net/http"
	"net/http/cookiejar"
	"os"
)

func CookieJar() {
	jar, _ := cookiejar.New(nil)
	client := &http.Client{Jar: jar}
	res, _ := http.NewRequest(http.MethodGet, "https://httpbin.org/cookies/set?name=xionger&age=18", nil)
	r, err := client.Do(res)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	_, _ = io.Copy(os.Stdout, r.Body)

}
