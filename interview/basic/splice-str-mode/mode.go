package splicestrmode

import (
	"bytes"
	"fmt"
	"strings"
)

func mode() {
	// 方式1  用 + 拼接
	s1 := "a"
	fmt.Println("方式1")
	fmt.Printf("%p\n", &s1)
	fmt.Println(s1)
	s1 = s1 + "b"
	fmt.Printf("%p\n", &s1) // s1地址一样，但是需要另外一个空间存a和ab
	fmt.Println(s1)

	// 方式2 用sprintf拼接，不常见跳过
	// 方式3 使用 strings.Builder
	fmt.Println("方式2")
	s2 := strings.Builder{}
	s2.WriteString("a")
	s2.WriteString("b")
	fmt.Println(s2.String())
	// 方式4 使用 bytes.Buffer
	fmt.Println("方式3")
	s3 := bytes.Buffer{}
	s3.WriteString("a")
	s3.WriteString("b")
	fmt.Println(s3.String())
	// 方式5 使用byte
	s4 := make([]byte, 0)
	s4 = append(s4, 'a')
	s4 = append(s4, 'b')
	fmt.Println(string(s4))
	// 如果长度已知，可以直接预先申请地址
	// 当然还有strings.join 底层还是builder
	// 应尽量选择strings.builder

}
