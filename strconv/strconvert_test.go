package strconv

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestBool2String(t *testing.T) {
	formatBool1 := strconv.FormatBool(true)
	assert.Equal(t, "true", formatBool1)
	formatBool2 := strconv.FormatBool(false)
	assert.Equal(t, "false", formatBool2)
}
