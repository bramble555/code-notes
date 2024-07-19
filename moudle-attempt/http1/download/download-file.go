package download

import (
	"fmt"
	"http1/basic"
	"io"
	"net/http"
	"os"
)

func Download() {
	url := "https://pic.baike.soso.com/ugc/baikepic2/10915/20220314112332-314610665_png_790_799_527527.jpg/0"
	path := "./download/file/xionger.png"
	r, err := http.Get(url)
	basic.EexamineErr(err)
	defer r.Body.Close()
	file, err := os.Create(path)
	basic.EexamineErr(err)
	io.Copy(file, r.Body)
	fmt.Println("下载完成")
}
