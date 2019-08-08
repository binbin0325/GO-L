package funcs

import (
	"fmt"
	"testing"
)

//当可变参数是一个空接口类型时，调用者是否解包可变参数会导致不同的结果：
func TestFunc(t *testing.T) {
	var a = []interface{}{123, "abc"}

	Print(a...) // 123 abc
	Print(a)    // [123 abc]
	fmt.Println(Inc())
	fmt.Println(f(2))
	fmt.Println(g())
}

// 可变数量的参数
// more 对应 []interface{} 切片类型
func Print(a ...interface{}) {
	fmt.Println(a...)
}

//返回值命名了，可以通过名字来修改返回值，也可以通过defer语句在return语句之后修改返回值：
func Inc() (v int) {
	defer func() { v++ }()
	return 42
}

func f(x int) *int {
	return &x
}

func g() int {
	x := new(int)
	return *x
}
