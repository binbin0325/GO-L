package const_test

import (
	"fmt"
	"testing"
)

const (
	a = iota
	b
	c
)

const (
	Unknown = "Unknown"
	Female  = "Female"
	Male    = "Male"
)

func TestConst(t *testing.T) {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(Unknown)
	fmt.Println(Female)
	fmt.Println(Male)
}
