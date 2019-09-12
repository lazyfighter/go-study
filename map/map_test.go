package _map

import (
	"log"
	"reflect"
	"testing"
)

func TestMapPut(t *testing.T) {
	m := make(map[string]interface{}, 10)
	m["test"] = "test"
	log.Println(len(m))
}

func TestDeepEqual(t *testing.T) {
	m := make(map[string]interface{}, 10)
	m["test"] = "test"

	m1 := make(map[string]interface{}, 10)
	m1["test"] = "test"
	log.Println(reflect.DeepEqual(m, m1))
}
