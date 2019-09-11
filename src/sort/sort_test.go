package sort

import (
	"log"
	"sort"
	"testing"
)

func TestSortInt(t *testing.T) {
	intArray := []int{2, 1, 3, 5, 4}
	sort.Ints(intArray)
	log.Print(intArray)
}

func TestSortFloat64(t *testing.T) {
	floatList := []float64{4.2, 5.9, 12.3, 10.0, 50.4, 99.9, 31.4, 27.81828, 3.14}
	sort.Float64s(floatList)
	log.Print(floatList)
}

func TestSortString(t *testing.T) {
	stringList := []string{"a", "c", "b", "d", "f", "i", "z", "x", "w", "y"}
	sort.Strings(stringList)
	log.Print(stringList)
}
