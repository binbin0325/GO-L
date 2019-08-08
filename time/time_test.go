package time_test

import (
	"fmt"
	"testing"
	"time"
)

var week time.Duration

func TestTime(t *testing.T) {
	t1 := time.Now()
	fmt.Println(t1)                                                //打印当前时间
	fmt.Printf("%02d-%02d-%4d\n", t1.Day(), t1.Month(), t1.Year()) //格式化输出时间

	t1 = time.Now().UTC() //UTC 表示通用协调世界时间。
	fmt.Println(t1)
	fmt.Println(time.Now())

	week = 60 * 60 * 24 * 7 * 1e9
	week_from_now := t1.Add(time.Duration(week)) //加7天
	fmt.Println(week_from_now)

	fmt.Println(t1.Format(time.RFC822))
	fmt.Println(t1.Format(time.ANSIC))
	fmt.Println(t1.Format("02 Jan 2006 15:04"))
	s := t1.Format("20060102")
	fmt.Println(t1, "=>", s)
}
