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

func TestDefer(t *testing.T) {
	defer func() {
		log.Println("defer")
	}()

	log.Println("start")
}
