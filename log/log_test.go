package log

import (
	"fmt"
	"log"
	"os"
	"testing"
)

// 测试新建log ， 三个参数，第一个输出需要实现io.Writer接口，第二个输出内容前缀，第三个flag 日志格式 在log文件中, 默认的flag
/*
	Ldate         = 1 << iota     // the date in the local time zone: 2009/01/23
	Ltime                         // the time in the local time zone: 01:23:23
	Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	Llongfile                     // full file name and line number: /a/b/c/d.go:23
	Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
	LUTC                          // if Ldate or Ltime is set, use UTC rather than the local time zone
	LstdFlags     = Ldate | Ltime // initial values for the standard logger
*/
func TestNewLog(t *testing.T) {

	logger := log.New(os.Stdout, "", log.Llongfile|log.LstdFlags)
	logger.Println("test stdout")

	file, e := os.Create("test.log")
	defer file.Close()
	if e != nil {
		fmt.Println("err", e)
	}

	//logger.SetOutput(file)

	//logger.Println("file out-put")

	// 可以打印调用栈信息
	logger.Output(1, "test")
}
