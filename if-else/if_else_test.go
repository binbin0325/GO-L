package if_else

import (
	"fmt"
	"runtime"
	"testing"
)

func TestIfElse(t *testing.T) {

}

var prompt = "Enter a digit,e.g.3" + "or %3 to quit."

//判断当前操作系统
func init() {
	if runtime.GOOS == "windows" {
		prompt = fmt.Sprintf(prompt, "Ctrl+Z,Enter")
	} else {
		prompt = fmt.Sprintf(prompt, "Ctrl+D")
	}
}

//绝对值
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

//比较两个整数类型的大小
func isGreater(x, y int) bool {
	if x > y {
		return true
	}
	return false
}
