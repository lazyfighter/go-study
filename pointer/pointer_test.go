package main

import (
	"log"
	"testing"
)

type S struct {
	data string
}

func (s S) Read() string {
	return s.data
}

func (s *S) Write(str string) {
	s.data = str
}

func TestPointer(t *testing.T) {
	log.Print("-------------")
	sVals := map[int]S{1: {"A"}}
	// 值类型变量只能调用 Read 方法
	sVals[1].Read()
	s := S{
		data: "11",
	}
	s.Write("dsad")

	sVals[2] = s

	//
	//// 无法编译通过:
	sVals[1].Write("test")
	sVals[2].Write("test")
	sPtrs := map[int]*S{1: {"A"}}
	// 指针类型变量可以调用 Read 和 Write 方法：
	log.Print(sPtrs[1].Read())
	sPtrs[1].Write("test")
}
