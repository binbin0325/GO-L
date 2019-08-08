package pointer_test

import (
	"fmt"
	"testing"
)

/*
*一个指针变量可以指向任何一个值的内存地址 它指向那个值的内存地址，
在 32 位机器上占用 4 个字节，在 64 位机器上占用 8 个字节，并且与它所指向的值的大小无关。
当然，可以声明指针指向任何类型的值来表明它的原始性或结构性；你可以在指针类型前面加上 * 号（前缀）来获取指针所指向的内容，
这里的 * 号是一个类型更改器。使用一个指针引用一个值被称为间接引用
*/

//声明i1=5,输出值和内存地址，声明intP 为指针类型，将intP = &i1 指向 i1的内存地址。
func TestPointer(t *testing.T) {
	var i1 = 5
	fmt.Printf("An integer : %d,it's location in memory: %p\n", i1, &i1) //&取指符

	var intP *int
	intP = &i1

	fmt.Printf("An integer : %d,it's location in memory: %p\n", *intP, intP)

}

func TestPointer1(t *testing.T) {
	s := "good bye"                               //声明s:="good bye" 声明p为指针类型
	var p *string = &s                            //将p指针指向s的内存地址
	*p = "ciao"                                   //将p的值修改为ciao,此时 s的值也是ciao了，因为p和s指向同一块内存地址，所以修改p  ，s也会改变
	fmt.Printf("Here is the pointer p: %p\n", p)  // prints address
	fmt.Printf("Here is the string *p: %s\n", *p) // prints string
	fmt.Printf("Here is the string s: %s\n", s)   // prints same string
}

//在Go中 是不能获取到一个文字或者常量的地址 例如：
//const i = 5
//ptr := &i   //error: cannot take the address of i
//ptr2 := &10 //error: cannot take the address of 10
