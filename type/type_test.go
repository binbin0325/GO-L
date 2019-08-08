package type_test

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"
)

//类型转换，go中没有隐式类型转换，类型转换必须显示的进行
func TestType(t *testing.T) {
	var n int16 = 34
	var m int32
	m = int32(n)
	fmt.Printf("32 bit int is : %d\n", m)
	fmt.Printf("16 bit int is : %d\n", n)
}

//安全的从int型转换为int8
func Uint8FromInt(n int) (uint8, error) {
	if 0 <= n && n <= math.MaxUint8 {
		return uint8(n), nil
	}
	return 0, fmt.Errorf("%d is out of the uint8 range", n)
}

//安全的从float64 转换为int
func IntFromFloat64(x float64) int {
	if math.MinInt32 <= x && x <= math.MaxInt32 {
		whole, fraction := math.Modf(x)
		if fraction >= 0.5 {
			whole++
		}
		return int(whole)
	}
	panic(fmt.Sprintf("%g is out of the int32 range", x))
}

//生成随机数
func TestRand(t *testing.T) {
	//rand,Int() 生成非负伪随机数
	for i := 0; i < 10; i++ {
		a := rand.Int()
		fmt.Printf("%d /", a)
	}
	//rand.Intn() 生成[0,n]之间的伪随机数
	for i := 0; i < 5; i++ {
		r := rand.Intn(8)
		fmt.Printf("%d /", r)
	}
	fmt.Println()
	//函数rand.Float32 和 rand.Float64 返回介于[0.0,1.0]之间的伪随机数，其中包括0.0 但不包括1.0，可以使用Seed(value)函数来提供伪随机数的生成种子，一般情况下会使用当前时间的纳秒级数字
	timens := int64(time.Now().Nanosecond())
	rand.Seed(timens)
	for i := 0; i < 10; i++ {
		fmt.Printf("%2.2f /", 100*rand.Float32())
	}
}
