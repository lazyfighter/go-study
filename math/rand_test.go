package math

import (
	"log"
	"math/rand"
	"testing"
)

func TestRand(t *testing.T) {
	rand := rand.New(rand.NewSource(1))
	for i := 0; i < 100; i++ {
		log.Println(rand.Intn(10))
	}
}
