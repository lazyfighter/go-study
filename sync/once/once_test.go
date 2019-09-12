package once

import (
	"log"
	"sync"
	"testing"
)

// 确保方法只执行一次，底层采用互斥锁，以及原子set,
func TestOnce(t *testing.T) {
	once := sync.Once{}
	for i := 1; i < 10; i++ {
		once.Do(func() {
			log.Print("once do")
		})
		log.Print("done")
	}
}
