package list

import (
	"container/list"
	"log"
	"testing"
)

func TestNew(t *testing.T) {
	list := list.New()
	list.PushFront("fuck")
}

type listElementTest struct {
	root list.Element
	len int
}

func TestNewMy(t *testing.T)  {
	test := new(listElementTest)
	log.Print(test.root)
}