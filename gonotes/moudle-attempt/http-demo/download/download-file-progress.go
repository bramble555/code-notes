package download

import (
	"fmt"
	"http1/basic"
	"io"
	"net/http"
	"os"
)

type ReaderProgress struct {
	io.Reader
	current int64
	total   int64
}

func sizeFormat(n int64) string {
	const (
		KB = 1024
		MB = KB * 1024
		GB = MB * 1024
	)
	switch {
	case n > GB:
		return fmt.Sprintf("%.2fGB", float64(n/GB))
	case n > MB:
		return fmt.Sprintf("%.2fMB", float64(n/MB))
	case n > KB:
		return fmt.Sprintf("%.2fKB", float64(n/KB))
	default:
		return fmt.Sprintf("%.f2B", float64(n))
	}
}

// ReaderProgress结构体实现Read方法
func (rp *ReaderProgress) Read(p []byte) (n int, err error) {
	n, err = rp.Reader.Read(p)
	rp.current += int64(n)
	// \r 的作用是让光标回到行首，这样下一次打印进度时就会覆盖掉上一次的信息，从而创建了一个动态更新的进度条效果。
	fmt.Printf("\r当前进度为%.2f%%=========文件大小为%s", float64(rp.current/rp.total)*100, sizeFormat(rp.total))
	// 一次read结束
	return
}
func DownloadProgress() {
	url := "https://pic.baike.soso.com/ugc/baikepic2/10915/20220314112332-314610665_png_790_799_527527.jpg/0"
	path := "./download/file/xionger2.png"
	r, err := http.Get(url)
	basic.EexamineErr(err)
	basic.EexamineErr(err)
	defer r.Body.Close()
	file, err := os.Create(path)
	basic.EexamineErr(err)
	readpro := &ReaderProgress{
		Reader: r.Body,
		total:  r.ContentLength,
	}
	_, err = io.Copy(file, readpro)
	basic.EexamineErr(err)
	if readpro.current == readpro.total {
		fmt.Println()
		fmt.Println("下载完成")
	}

}
