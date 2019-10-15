package slice

import (
	"log"
	"reflect"
	"testing"
)

func TestSlice(t *testing.T) {
	stringSlice := make([]int, 0, 3)
	log.Print(stringSlice, len(stringSlice), cap(stringSlice))
	stringSlice = append(stringSlice, 1)
	log.Print(stringSlice, len(stringSlice), cap(stringSlice))
}

func TestSlice1(t *testing.T) {
	stringSlice := make([]int, 3, 3)
	log.Print(reflect.TypeOf(stringSlice).Name(), reflect.TypeOf(stringSlice).Kind())
	log.Print(stringSlice, len(stringSlice), cap(stringSlice))
	stringSlice = append(stringSlice, 1)
	stringSlice[0] = 2
	log.Print(stringSlice, len(stringSlice), cap(stringSlice))
}

// 数组不能够使用append
func TestArray(t *testing.T) {
	var names [5]string
	names[0] = "li"
	log.Print(names)
}

func TestName(t *testing.T) {

}
