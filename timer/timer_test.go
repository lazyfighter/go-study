package timer

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// 创建一个timer
func TestTimer(t *testing.T) {
	timer := time.NewTimer(10 * time.Second)
	fmt.Println("create timer", time.Now().Format("2006-01-02 15:04:05"))
	<-timer.C
	fmt.Println("timer done", time.Now().Format("2006-01-02 15:04:05"))
}

// 对timer进行reset操作，复用原先的timer的数据，但是修改了duration，然后从新启动一个timer
func TestTimerReset(t *testing.T) {
	timer := time.NewTimer(10 * time.Second)
	fmt.Println("create timer", time.Now().Format("2006-01-02 15:04:05"))
	<-timer.C
	fmt.Println("timer done", time.Now().Format("2006-01-02 15:04:05"))
	timer.Reset(1 * time.Second)
	<-timer.C
	fmt.Println("timer done", time.Now().Format("2006-01-02 15:04:05"))
}

// timer 可以进行stop操作
func TestTimerStop(t *testing.T) {
	timer := time.NewTimer(1 * time.Second)
	fmt.Println("create timer", time.Now().Format("2006-01-02 15:04:05"))
	timer.Stop()
	<-timer.C
	fmt.Println("timer done", time.Now().Format("2006-01-02 15:04:05"))
}

// 返回的是个通道，底层依然使用newTimer 但是返回的通道无法进行stop操作
func TestAfter(t *testing.T) {
	after := time.After(10 * time.Second)
	fmt.Println("create timer", time.Now().Format("2006-01-02 15:04:05"))
	<-after
	fmt.Println("timer done", time.Now().Format("2006-01-02 15:04:05"))

}

// 返回的是个通道，底层依然使用newTimer 但是返回的通道无法进行stop操作
func TestAfterFunc(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	time.AfterFunc(1*time.Second, func() {
		fmt.Println("timer done", time.Now().Format("2006-01-02 15:04:05"))
		wg.Done()
	})
	wg.Wait()
}

// 与timer进行对比，ticker可以执行多次
func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)
	fmt.Println("create timer", time.Now().Format("2006-01-02 15:04:05"))
	count := 10
	for count > 0 {
		<-ticker.C
		fmt.Println("timer done", time.Now().Format("2006-01-02 15:04:05"))
		count--
	}
}

// tick 可以进行停止，但是依然从tick进行通道获取数据，会导致死锁， timer同样
func TestTickerStop(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)
	fmt.Println("create timer", time.Now().Format("2006-01-02 15:04:05"))
	count := 10
	for count > 0 {
		<-ticker.C
		fmt.Println("timer done", time.Now().Format("2006-01-02 15:04:05"), count)
		if count == 3 {
			ticker.Stop()
		}
		count--
	}
}

// tick 没有提供stop 功能，类似于 time.After 两者的区别同样tick可以执行多次知道stop ， after只能执行一次
func TestTick(t *testing.T) {
	ticker := time.Tick(1 * time.Second)
	fmt.Println("create timer", time.Now().Format("2006-01-02 15:04:05"))
	count := 10
	for count > 0 {
		<-ticker
		fmt.Println("timer done", time.Now().Format("2006-01-02 15:04:05"), count)
		count--
	}
}

func TestNanoTime(t *testing.T) {
	fmt.Println(time.Now().UnixNano())
	fmt.Println(int64(time.Now().UnixNano()) / int64(time.Millisecond))
}
