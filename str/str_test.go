package str_test

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"unicode/utf8"
	"unsafe"
)

/*Go中的字符串里面的字符也可能根据需要占用1至4个字节，（java始终使用2个字节），
Go这样做的好处是不仅减少了内存和硬盘空间占用，同时也不用像其它语言那样需要对使用UTF-8字符集的文本进行编码和解码

字符串是一种值类型，且值不可变，即创建某个文本后 你无法再次修改这个文本的内容，更深入的讲，字符串是字节的定长数组
*/
func TestCountCharacters(t *testing.T) {
	str1 := "asSASA ddd dsjkdsjs dk"
	fmt.Printf("The number of bytes in string str1 is %d\n", len(str1))
	fmt.Printf("The number of characters in string str1 is %d\n", utf8.RuneCountInString(str1))
	str2 := "asSASA ddd dsjkdsjsこん dk"
	fmt.Printf("The number of bytes in string str2 is %d\n", len(str2))
	fmt.Printf("The number of characters in string str2 is %d", utf8.RuneCountInString(str2))
}

//利用strings.HasPrefix 判断字符串是否已Th开头 以ing结尾，区分大小写
func TestFun1(t *testing.T) {
	str := "This is an example of a string"
	fmt.Printf("T/F? Does the string \"%s\" have prefix %s?", str, "th")
	fmt.Printf("\n%t", strings.HasPrefix(str, "Th"))
	fmt.Printf("\n%t", strings.HasSuffix(str, "ing"))
}

//判断字符串是否包含指定字符   strings.Contains()
func TestContains(t *testing.T) {
	str := "This is an example of a string"
	fmt.Printf("\n%t", strings.Contains(str, "an e"))
}

//判断字符出现的位置  strings.Index()
func TestIndex(t *testing.T) {
	str := "This is an example of a string"
	fmt.Printf("\n%d", strings.Index(str, "an"))

	fmt.Printf("\n%d", strings.LastIndex(str, "a"))
}

//字符串替换
func TestReplace(t *testing.T) {
	str := "This is an example of a string"
	fmt.Print(strings.Replace(str, "an", "xx", -1))
}

//统计字符串出现次数
func TestCount(t *testing.T) {
	str := "This is an example of a string"
	fmt.Print(strings.Count(str, "g"))
}

//重复字符串
func TestRepeat(t *testing.T) {
	origs := "Hi there! "
	var newS string
	newS = strings.Repeat(origs, 3) //3 是重复次数
	fmt.Printf("%s\n", newS)
}

//修改字符串大小写
func TestTo(t *testing.T) {
	str := "Hi"
	fmt.Println(strings.ToLower(str)) //小写
	fmt.Println(strings.ToUpper(str)) //大写
}

//修剪字符串
func TestTrimSpace(t *testing.T) {
	str := " This is an example of a string"
	fmt.Println(strings.TrimSpace(str))      //剔除字符串开头和结尾的空白符号
	fmt.Println(strings.Trim(str, "string")) //剔除开头和结尾的任何字符
	//如果只想剔除开头和结尾的字符串，则可以使用TrimLeft 或者 TrimRight来实现
}

//分割字符串
func TestSplic(t *testing.T) {
	str := "This is an example of a string"
	fields := strings.Fields(str) //利用空白符号来作为动态长度的分隔符将字符串分隔成若干小块，并返回一个slice
	for i := 0; i < len(fields); i++ {
		fmt.Println(fields[i])
	}
	fmt.Println("******************************")
	splits := strings.Split(str, " ")
	for i := 0; i < len(splits); i++ {
		fmt.Println(splits[i])
	}

}

//拼接slice 到字符串
func TestJoin(t *testing.T) {
	str := "This is an example of a string"
	fields := strings.Fields(str)
	for _, val := range fields {
		fmt.Printf("%s - ", val)
	}
}

//类型转换
func TestConv(t *testing.T) {
	//数字类型转换到字符串类型
	strconv.Itoa(100)
	//将浮点类型转换为字符串 ---
	//strconv.FormatFloat(f float64, fmt byte, prec int, bitSize int) string 将 64 位浮点型的数字转换为字符串，其中 fmt 表示格式（其值可以是 'b'、'e'、'f' 或 'g'），prec 表示精度，bitSize 则使用 32 表示 float32，用 64 表示 float64。
	fmt.Println(strconv.FormatFloat(23.983, 'f', 2, 64))

	//将字符串转换为数字类型
	strconv.Atoi("123")
	fmt.Println(strconv.ParseFloat("123", 64))
}

/**
截取字符串[:5] 截取前五位，[5:] 从第6位开始截取
*/
func TestS1(t *testing.T) {
	s := "hello, world"
	hello := s[:5]
	world := s[7:]
	fmt.Println(hello)
	fmt.Println(world)
	s1 := "hello, world"[:5]
	s2 := "hello, world"[7:]
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println("len(s):", (*reflect.StringHeader)(unsafe.Pointer(&s)).Len)   // 12
	fmt.Println("len(s1):", (*reflect.StringHeader)(unsafe.Pointer(&s1)).Len) // 5
	fmt.Println("len(s2):", (*reflect.StringHeader)(unsafe.Pointer(&s2)).Len) // 5

	//Go语言除了for range语法对UTF8字符串提供了特殊支持外，还对字符串和[]rune类型的相互转换提供了特殊的支持。
	fmt.Printf("%#v\n", []rune("世界"))             // []int32{19990, 30028}
	fmt.Printf("%#v\n", string([]rune{'世', '界'})) // 世界

	g := make([]string, 3)
	g = append(g, "你好")
	g = append(g, " ")
	g = append(g, "golang")
	fmt.Println(g)
	fmt.Println(TrimSpace(g))
}

//TrimSpace函数用于删除[]string中的空格。函数实现利用了0长切片的特性，实现高效而且简洁。
func TrimSpace(s []string) []string {
	b := s[:0]
	for _, x := range s {
		if x != " " {
			b = append(b, x)
		}
	}
	return b
}
