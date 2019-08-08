package var_test

import (
	"fmt"
	"math"
	"os"
	"runtime"
	"testing"
)

type Student struct {
	name string
	sex  int
}

var NAME string

func TestVar(t *testing.T) {
	var a, b int
	a = 1
	b = 2
	fmt.Println(a)
	fmt.Println(b)

	NAME = "张三"
	NAME := "张四"
	fmt.Println(NAME)

	c := []int{1, 2, 3, 4}
	d := c[:2]
	d[1] = 1
	fmt.Println(cap(d))
	fmt.Println(c)
	fmt.Println(d)
}

//声明变量    =赋值 和:= 赋值的区别， :=赋值 是针对于变量初始化赋值，:=会创建一个新的变量，作用域在函数局部
func TestVar1(t *testing.T) {
	var goos = runtime.GOOS
	fmt.Printf("The operating system is: %s\n", goos)
	path := os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)
}

//值拷贝
func TestVar2(t *testing.T) {
	i := 7
	j := i
	fmt.Println(&i)
	fmt.Println(&j)
}

func TestVar3(t *testing.T) {
	fmt.Println(Pi)
}

//交换两个变量的值
func TestVar4(t *testing.T) {
	a := 1
	b := 2
	a, b = b, a
	fmt.Print(a, b)
}

var Pi float64

//init函数 执行优先级比main函数高，每个源文件都只能包含一个init函数，初始化总是以单线程执行，并且按照包的依赖关系顺序执行
func init() {
	Pi = 4 * math.Atan(1)
}

var a string

func TestVar5(t *testing.T) {
	a = "G"
	print(a)
	n()
}
func n() {
	a := "O"
	print(a)
	m()
}
func m() {
	print(a)
}
