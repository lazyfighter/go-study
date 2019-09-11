package timer

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// 有单位的时间字符串，time提供ParseDuration能够将时间转为Duration ， 适合配置超时时间这种很方便
func TestParseDuration(t *testing.T) {
	duration, err := time.ParseDuration("2s")
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println("duration", duration.Hours())
	fmt.Println("duration", duration.Seconds())
	group := sync.WaitGroup{}
	group.Add(1)
	fmt.Println("timer create", time.Now().Format("2006-01-02 15:04:05"))
	time.AfterFunc(duration, func() {
		fmt.Println("timer done", time.Now().Format("2006-01-02 15:04:05"))
		group.Done()
	})
	group.Wait()

}

// 自定义模板进行时间转换，很灵活, 注意默认使用的UTC时区
func TestParse(t *testing.T) {

	format, e := time.Parse("2006-01-02", "2010-02-04")
	if e != nil {
		fmt.Println("err", e)
		return
	}
	fmt.Println(format.String())
	format, e = time.Parse("2006-01-02 15:04:05", "2010-02-04 21:00:57.0")
	if e != nil {
		fmt.Println("err", e)
		return
	}
	fmt.Println(format.String())
}

// 指定时区进行时间的转换
func TestParseInLocation(t *testing.T) {
	format, e := time.ParseInLocation("2006-01-02 15:04:05", "2010-02-04 21:00:57.0", time.FixedZone("test", 7))
	if e != nil {
		fmt.Println("err", e)
		return
	}
	fmt.Println(format)

	// 使用本机的时区
	format, e = time.ParseInLocation("2006-01-02 15:04:05", "2010-02-04 21:00:57.0", time.Local)
}

// 按照指定的模板进行时间格式化，此处需要注意，模板的时间必须是这个时间，go语言诞生的时间， 默认使用本地时区
func TestFormat(t *testing.T) {
	format := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(format)
}

// 按照指定的模板进行时间格式化，同时append到指定的数组中
func TestAppendFormat(t *testing.T) {
	date := time.Date(2017, 12, 4, 11, 0, 0, 0, time.UTC)
	text := []byte("Time: ")
	text = date.AppendFormat(text, "2006-01-02 15:04:05")
	fmt.Println(string(text))
}
