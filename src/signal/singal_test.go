package signal

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"
)

// Notify 提供两个参数 ，第一个参数为接收信号的管道， 第二个参数为设置要监听的信号，如果不设置则表示监听所有的信号
func TestSignalNotify(t *testing.T) {
	chanSignal := make(chan os.Signal, 1)

	signal.Notify(chanSignal, os.Kill)

	go func() {
		time.Sleep(1 * time.Second)
		chanSignal <- os.Interrupt
	}()

	ss := <-chanSignal

	fmt.Println("rec", ss)

}

// 当调用stop方法之后，只要接收到一个信号，使用到管道的地方就会推出
func TestSignalNotifyStop(t *testing.T) {
	chanSignal := make(chan os.Signal, 1)

	signal.Notify(chanSignal)

	fmt.Println("pid", os.Getpid())
	count := 10
	for {
		ss := <-chanSignal
		if count < 0 {
			signal.Stop(chanSignal)
		}
		count--
		fmt.Println("rec", ss, count)
	}
}

// 忽略信号
func TestIgnore(t *testing.T) {
	fmt.Println("pid", os.Getpid())
	chanSignal := make(chan os.Signal, 1)

	signal.Notify(chanSignal)

	signal.Ignore(syscall.SIGBUS)

	ss := <-chanSignal
	fmt.Println("rec", ss)

}

// 判断这个信号是否被忽略
func TestIgnored(t *testing.T) {
	fmt.Println("pid", os.Getpid())
	chanSignal := make(chan os.Signal, 1)

	signal.Notify(chanSignal)

	signal.Ignore(syscall.SIGBUS)

	ignored := signal.Ignored(syscall.SIGBUS)

	fmt.Println("ignored", ignored)

	ignored = signal.Ignored(os.Kill)

	fmt.Println("ignored", ignored)

}

// 撤销之之前指定程序能接收的信号，但是可以重新使用Notify指定接收的信号。
func TestReset(t *testing.T) {
	chanSignal := make(chan os.Signal, 1)

	signal.Notify(chanSignal)

	ignored := signal.Ignored(os.Kill)

	fmt.Println("ignored", ignored)

	signal.Reset(os.Interrupt)

	ignored = signal.Ignored(os.Kill)

	fmt.Println("ignored", ignored)

}
