package reflect

import (
	"log"
	"reflect"
	"strings"
	"testing"
)

// 反射获取类型的名称
func TestName(t *testing.T) {
	of := reflect.TypeOf(person{})
	log.Println(of.Name())
}

// 反射类型的toString
func TestString(t *testing.T) {
	of := reflect.TypeOf(person{})
	log.Println(of.String())
}

// 获取map的key的类型
func TestKey(t *testing.T) {
	of := reflect.TypeOf(person{}.mapTest)
	log.Println(of.Key())
}

// 获取对象的方法数量
func TestNumMethod(t *testing.T) {
	of := reflect.TypeOf(&person{})
	log.Println(of.NumMethod())
}

// 获取对象的方法相关,然后通过发射调用方法
func TestMethod(t *testing.T) {
	person := &person{}
	of := reflect.TypeOf(person)
	for i := 0; i < of.NumMethod(); i++ {
		method := of.Method(i)
		log.Print("name : ", method.Name, ", index: ", method.Index, ", pakPath:", method.PkgPath, ", type: ", method, " numIn:", method.Type.NumIn())
		fun := method.Func
		if strings.HasPrefix(method.Name, "Set") {

			// 判断方法参数是否为... 不固定数量类型
			if fun.Type().IsVariadic() {
				fun.CallSlice([]reflect.Value{reflect.ValueOf(person), reflect.ValueOf([]string{"basketball", "pingPang"})})
			} else {
				if fun.Type().In(1).AssignableTo(reflect.TypeOf(int(0))) {
					fun.Call([]reflect.Value{reflect.ValueOf(person), reflect.ValueOf(int(32))})
				} else {
					fun.Call([]reflect.Value{reflect.ValueOf(person), reflect.ValueOf("testName")})
				}
			}
		}
	}
	log.Print(person.name, person.age, person.hobbies)
}
