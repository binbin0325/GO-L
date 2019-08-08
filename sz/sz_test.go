package sz

import (
	"fmt"
	"testing"
)

func TestM(t *testing.T) {
	var times [5][1]int
	for i, v := range times {
		fmt.Println(i, ":")
		for range v {
			fmt.Println("Hello")
		}
	}
}

func TestM1(t *testing.T) {
	var a [3]int                    // 定义长度为3的int型数组, 元素全部为0
	var b = [...]int{1, 2, 3}       // 定义长度为3的int型数组, 元素为 1, 2, 3
	var c = [...]int{2: 3, 1: 2}    // 定义长度为3的int型数组, 元素为 0, 2, 3
	var d = [...]int{1, 2, 4: 5, 6} // 定义长度为6的int型数组, 元素为 1, 2, 0, 0, 5, 6
	fmt.Println(a, b, c, d)
}

/*
 *当一个数组变量被赋值或者被传递的时候，实际上会复制整个数组。
 *如果数组较大的话，数组的赋值也会有较大的开销。
 *为了避免复制数组带来的开销，可以传递一个指向数组的指针，但是数组指针并不是数组
 */
func TestM2(t *testing.T) {
	var a = [...]int{1, 2, 3} // a 是一个数组
	var b = &a                // b 是指向数组的指针
	fmt.Println(a[0], a[1])   // 打印数组的前2个元素
	fmt.Println(b[0], b[1])   // 通过数组指针访问数组元素的方式和数组类似
	//用for range方式迭代的性能可能会更好一些，因为这种迭代可以保证不会出现数组越界的情形，每轮迭代对数组元素的访问时可以省去对下标越界的判断。
	for i, v := range b { // 通过数组指针迭代数组的元素
		fmt.Println(i, v)
	}
}
