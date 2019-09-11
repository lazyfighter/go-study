package reflect

import (
	"log"
	"reflect"
	"testing"
)

func TestDeepEqual(t *testing.T) {

	p1 := person{name: "test", age: 18, hobbies: []string{"aa", "bb"}}
	p2 := person{name: "test", age: 18, hobbies: []string{"aa", "bb"}}

	log.Println(reflect.DeepEqual(p1, p2))
}

func TestErrorEqual(t *testing.T) {
	m1 := map[string]string{"a": "1", "b": "2", "c": "3"}
	m2 := map[string]string{"a": "1", "c": "3", "b": "2"}
	log.Println(cmpMap(m1, m2))
	log.Println(reflect.DeepEqual(m1, m2))
}

func cmpMap(m1, m2 map[string]string) bool {
	for k1, v1 := range m1 {
		if v2, has := m2[k1]; has {
			if v1 != v2 {
				return false
			}
		} else {
			return false
		}
	}
	for k2, v2 := range m2 {
		if v1, has := m1[k2]; has {
			if v1 != v2 {
				return false
			}
		} else {
			return false
		}
	}
	return true
}
