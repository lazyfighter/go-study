package strconv

import (
	"github.com/stretchr/testify/assert"
	"log"
	"reflect"
	"strconv"
	"testing"
)

func TestBool2String(t *testing.T) {
	formatBool1 := strconv.FormatBool(true)
	assert.Equal(t, "true", formatBool1)
	formatBool2 := strconv.FormatBool(false)
	assert.Equal(t, "false", formatBool2)
}

func TestParseFloat64(t *testing.T) {
	f, e := strconv.ParseFloat("100.10", 64)
	if e != nil {
		log.Fatal(e)
	}
	log.Print(f, reflect.TypeOf(f).Name())
	assert.Equal(t, float64(100.10), f)

}

// 虽然转换成为了float32 但是类型还是float64
func TestParseFloat32(t *testing.T) {
	f, e := strconv.ParseFloat("100.10", 32)
	if e != nil {
		log.Fatal(e)
	}
	log.Print(f, reflect.TypeOf(f).Name())

}

func TestFormatFloat64(t *testing.T) {
	formatFloat := strconv.FormatFloat(float64(100.10), 'e', 2, 64)
	log.Print(formatFloat)
}
