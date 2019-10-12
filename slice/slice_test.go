package slice

import (
	"log"
	"testing"
)

func TestSlice(t *testing.T) {
	stringSlice := make([]int, 0, 3)
	log.Print(stringSlice, len(stringSlice), cap(stringSlice))
	stringSlice = append(stringSlice, 1)
	log.Print(stringSlice, len(stringSlice), cap(stringSlice))
}
