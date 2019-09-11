package builtin

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// panic goroutine
func TestPanic(t *testing.T) {

	group := sync.WaitGroup{}
	group.Add(2)

	go func() {
		fmt.Println("fuck 1 go ")
		panic("panic")
		group.Done()
	}()

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("fuck 2 go ")
		group.Done()
	}()

	group.Wait()
}

// 使用recover 来捕捉异常
func TestRecover(t *testing.T) {
	group := sync.WaitGroup{}
	group.Add(2)

	go func() {
		defer func() {
			fmt.Println("defer func")
			if err := recover(); err != nil {
				fmt.Println("recover error", err)
			}
			group.Done()
		}()
		fmt.Println("fuck 1 go ")
		panic("panic")
	}()

	go func() {

		defer group.Done()

		time.Sleep(2 * time.Second)
		fmt.Println("fuck 2 go ")

	}()

	group.Wait()
}
