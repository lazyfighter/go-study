package reflect

import (
	"log"
	"reflect"
	"testing"
)

// 通过反射获取类型
func TestReflectStructType(t *testing.T) {
	rt := reflect.TypeOf(person{})
	log.Println(rt.Kind())
}

func TestReflectMapType(t *testing.T) {
	i := make(map[string]interface{})
	i["test"] = "testValue"
	i["test1"] = 2
	i["test2"] = nil
	rt := reflect.TypeOf(i)
	log.Println(rt.Kind())
	log.Println(rt.Key())

	v := reflect.ValueOf(i)
	for _, key := range v.MapKeys() {
		log.Print(v.MapIndex(key).CanInterface())
	}
}

func TestReflectNilType(t *testing.T) {
	rt := reflect.TypeOf(nil)
	log.Println(rt)
}

func TestReflectNumberType(t *testing.T) {
	var ss = 10.0
	rt := reflect.TypeOf(ss)
	log.Println(rt)
}
