package path

import (
	"fmt"
	"log"
	"path"
	"path/filepath"
	"testing"
)

func TestDir(t *testing.T) {
	dir := path.Dir("/data/test/a/b")
	fmt.Println("dir:", dir)
}

func TestExt(t *testing.T) {
	ext := path.Ext("/test/a.txt")
	fmt.Println("ext:", ext)
}

func TestGetCurrentDir(t *testing.T) {
	dir, _ := filepath.Abs(".")
	log.Println(dir)
}
