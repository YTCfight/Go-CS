package test

import "fmt"

// 结构体名字和成员的首字母都要大写，外包才能访问
type Student struct {
	Name string
	Age  int
	Sex  string
}

func init() {
	fmt.Println("this is test.init")
}
