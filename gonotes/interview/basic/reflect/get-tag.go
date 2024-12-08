package reflect

import (
	"fmt"
	"reflect"
)

type Author struct {
	Name         string `json:"name"`
	Publications string `json:"publications"`
}

func get() {
	t := reflect.TypeOf(Author{})
	for i := 0; i < t.NumField(); i++ {
		name := t.Field(i).Name
		na, ok := t.FieldByName(name)
		fmt.Println(i, name, na, ok)
	}
	author := Author{"1", "2"}
	fmt.Printf("%v\n", author)  // %v输出结构体各成员的值
	fmt.Printf("%+v\n", author) // %+v输出结构体各成员的名称和值；
	fmt.Printf("%#v\n", author) // %#v输出结构体名称和结构体各成员的名称和值

}
