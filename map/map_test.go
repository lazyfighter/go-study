package _map

import (
	"log"
	"testing"
)

func TestMap(t *testing.T) {
	m := make(map[string]interface{}, 10)
	m["test"] = "test"
	log.Println(len(m))
}
