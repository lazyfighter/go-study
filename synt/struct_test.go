package synt

import (
	"log"
	"testing"
)

// 初始化空的struct
func TestStructSyntax(t *testing.T) {

	type s struct{}
	invalidRequest := s{}
	invalidRequest = struct{}{}
	log.Println(invalidRequest)

}

func TestConst(t *testing.T) {
	const aa = 11
	log.Println(aa)
}
