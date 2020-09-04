package models

import (
	"fmt"
	"testing"
	"unsafe"
)

// Go 内存对齐

type s1 struct {
	a int8   // 1
	b string // 3
	c int8   // 1
}

type s2 struct {
	a int8
	c int8   // 1
	b string // 2
}

func TestStruct(t *testing.T) {
	v1 := s1{
		a: 1,
		b: "七米",
		c: 2,
	}

	v2 := s2{
		a: 1,
		b: "七米",
		c: 2,
	}

	fmt.Println(unsafe.Sizeof(v1), unsafe.Sizeof(v2))
}
