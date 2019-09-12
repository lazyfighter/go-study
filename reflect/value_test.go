package reflect

import (
	"log"
	"reflect"
	"testing"
	"time"
)

var p = person{
	name:    "test",
	age:     18,
	hobbies: []string{"aa", "bb"},
}

// 如果参数为指针，返回指针指向的interface， 如果变量不是指针，返回变量本身
func TestIndirect(t *testing.T) {
	indirect := reflect.Indirect(reflect.ValueOf(&p))
	log.Println(indirect)
	log.Println(indirect.CanAddr())
	indirect = reflect.Indirect(reflect.ValueOf(p))
	log.Println(indirect)
	log.Println(indirect.CanAddr())

}

func TestCanAddr1(t *testing.T) {
	p.boy = &struct{}{}
	log.Println(reflect.ValueOf(&p).Elem().CanAddr())
	log.Println(reflect.ValueOf(&p).Elem().CanSet())
	log.Println(reflect.ValueOf(&p).Elem().CanAddr())
}

func TestAddr(t *testing.T) {
	//p.SetAge(20)
	//p.Test()
	// struct
	// Test 方法挂载到struct上面
	log.Println(reflect.ValueOf(p).MethodByName("Test").Call([]reflect.Value{}))

	log.Println(reflect.ValueOf(&p).Kind())
	log.Println(reflect.ValueOf(&p).Elem().Addr().Kind())
	//ptr 方法挂载在ptr上面
	log.Println(reflect.ValueOf(&p).Elem().Addr().MethodByName("SetAge").Call([]reflect.Value{reflect.ValueOf(int(20))}))
	log.Println(p.age)
}

type S struct {
	X int
	Y string
	z int
}

func M() int {
	return 100
}

// copy form https://colobu.com/2018/02/27/go-addressable/
func TestCanAddr2(t *testing.T) {
	// 基础变量不可以寻址
	var x0 = 0
	v := reflect.ValueOf(x0)
	log.Printf("x0: %v \tcan be addressable and set: %t, %t\n", x0, v.CanAddr(), v.CanSet()) //false,false

	// 还是基础变量寻址
	var x1 = 1
	v = reflect.Indirect(reflect.ValueOf(x1))
	log.Printf("x1: %v \tcan be addressable and set: %t, %t\n", x1, v.CanAddr(), v.CanSet()) //false,false

	// 对指针类型进行寻址，因此都为true
	var x2 = &x1
	v = reflect.Indirect(reflect.ValueOf(x2))
	log.Printf("x2: %v \tcan be addressable and set: %t, %t\n", x2, v.CanAddr(), v.CanSet()) //true,true

	// 对非指针进行寻址，还是原先的对象，都为false
	var x3 = time.Now()
	v = reflect.Indirect(reflect.ValueOf(x3))
	log.Printf("x3: %v \tcan be addressable and set: %t, %t\n", x3, v.CanAddr(), v.CanSet()) //false,false

	// 对指针进行寻找找到了指针对应的对象，可以寻址可以改变都为true
	var x4 = &x3
	v = reflect.Indirect(reflect.ValueOf(x4))
	log.Printf("x4: %v \tcan be addressable and set: %t, %t\n", x4, v.CanAddr(), v.CanSet()) // true,true

	// 对数组进行寻址，还是其本身，无法寻址不能改变，都为false
	var x5 = []int{1, 2, 3}
	v = reflect.ValueOf(x5)
	log.Printf("x5: %v \tcan be addressable and set: %t, %t\n", x5, v.CanAddr(), v.CanSet()) // false,false

	// 对数组里面的某个元素进行寻址，无法寻址不能改变，都为false
	var x6 = []int{1, 2, 3}
	v = reflect.ValueOf(x6[0])
	log.Printf("x6: %v \tcan be addressable and set: %t, %t\n", x6[0], v.CanAddr(), v.CanSet()) //false,false

	// 对数组里面的某个元素通过索引进行寻址，都为true
	var x7 = []int{1, 2, 3}
	v = reflect.ValueOf(x7).Index(0)
	log.Printf("x7: %v \tcan be addressable and set: %t, %t\n", x7[0], v.CanAddr(), v.CanSet()) //true,true

	// 对数组里面的某个元素通过指针进行寻址，都为true
	v = reflect.ValueOf(&x7[1])
	log.Printf("x7.1: %v \tcan be addressable and set: %t, %t\n", x7[1], v.CanAddr(), v.CanSet()) //true,true

	// 对数组里面的某个元素进行寻址，无法寻址不能改变，都为false
	var x8 = [3]int{1, 2, 3}
	v = reflect.ValueOf(x8[0])
	log.Printf("x8: %v \tcan be addressable and set: %t, %t\n", x8[0], v.CanAddr(), v.CanSet()) //false,false

	// 拿到的是值得copy，因此无法进行寻址
	var x9 = [3]int{1, 2, 3}
	v = reflect.Indirect(reflect.ValueOf(x9).Index(0))
	log.Printf("x9: %v \tcan be addressable and set: %t, %t\n", x9[0], v.CanAddr(), v.CanSet()) //false,false

	var x10 = [3]int{1, 2, 3}
	v = reflect.Indirect(reflect.ValueOf(&x10)).Index(0)
	log.Printf("x9: %v \tcan be addressable and set: %t, %t\n", x10[0], v.CanAddr(), v.CanSet()) //true,true

	var x11 = S{}
	v = reflect.ValueOf(x11)
	log.Printf("x11: %v \tcan be addressable and set: %t, %t\n", x11, v.CanAddr(), v.CanSet()) //false,false
	var x12 = S{}
	v = reflect.Indirect(reflect.ValueOf(&x12))
	log.Printf("x12: %v \tcan be addressable and set: %t, %t\n", x12, v.CanAddr(), v.CanSet()) //true,true
	var x13 = S{}
	v = reflect.ValueOf(x13).FieldByName("X")
	log.Printf("x13: %v \tcan be addressable and set: %t, %t\n", x13, v.CanAddr(), v.CanSet()) //false,false
	var x14 = S{}
	v = reflect.Indirect(reflect.ValueOf(&x14)).FieldByName("X")
	log.Printf("x14: %v \tcan be addressable and set: %t, %t\n", x14, v.CanAddr(), v.CanSet()) //true,true
	var x15 = S{}
	v = reflect.Indirect(reflect.ValueOf(&x15)).FieldByName("z")
	log.Printf("x15: %v \tcan be addressable and set: %t, %t\n", x15, v.CanAddr(), v.CanSet()) //true,false
	v = reflect.Indirect(reflect.ValueOf(&S{}))
	log.Printf("x15.1: %v \tcan be addressable and set: %t, %t\n", &S{}, v.CanAddr(), v.CanSet()) //true,true
	var x16 = M
	v = reflect.ValueOf(x16)
	log.Printf("x16: %p \tcan be addressable and set: %t, %t\n", x16, v.CanAddr(), v.CanSet()) //false,false
	var x17 = M
	v = reflect.Indirect(reflect.ValueOf(&x17))
	log.Printf("x17: %p \tcan be addressable and set: %t, %t\n", x17, v.CanAddr(), v.CanSet()) //true,true

	var x18 interface{} = &x11
	v = reflect.ValueOf(x18)
	log.Printf("x18: %v \tcan be addressable and set: %t, %t\n", x18, v.CanAddr(), v.CanSet()) //false,false
	var x19 interface{} = &x11
	v = reflect.ValueOf(x19).Elem()
	log.Printf("x19: %v \tcan be addressable and set: %t, %t\n", x19, v.CanAddr(), v.CanSet()) //true,true

	var x20 = [...]int{1, 2, 3}
	v = reflect.ValueOf([...]int{1, 2, 3})
	log.Printf("x20: %v \tcan be addressable and set: %t, %t\n", x20, v.CanAddr(), v.CanSet()) //false,false
}
