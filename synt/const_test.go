package synt

import (
	"log"
	"testing"
)

// 当地一个int类型的变量赋值为iota， 第二个会进行递增
const (
	test1 int = iota
	test2
)

const (
	test3 int = 3
	test4
)

func TestConst(t *testing.T) {
	log.Print(test1)
	log.Print(test2)
	log.Print(test3)
	log.Print(test4)
}
