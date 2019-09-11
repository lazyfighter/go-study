package buffer

import (
	"log"
	"testing"
)

func TestMaxInt(t *testing.T) {

	log.Printf("%T , c = %b \n", uint(0), uint(0))
	log.Printf("%T , c = %b \n", ^uint(0), ^uint(0))
	log.Printf("%T , c = %b \n", int(^uint(0)>>1), int(^uint(0)>>1))

}
