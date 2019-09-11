package strings

import (
	"log"
	"strconv"
	"testing"
)

func TestEqual(t *testing.T) {
	s1 := "aaa"
	s2 := "aaa"
	log.Print(s1 == s2)
	s3 := "bbb"
	log.Print(s1 == s3)
	s4 := string("aaa")
	log.Print(s1 == s4)

}

// 将整数转换为字符串
func TestItoa(t *testing.T) {
	itoa := strconv.Itoa(120)
	log.Println(itoa)
}
