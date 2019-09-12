package context

import (
	"context"
	"log"
	"testing"
	"time"
)

func TestContext(t *testing.T) {

	timeout, _ := context.WithTimeout(context.Background(), 1*time.Second)
	time.Sleep(2 * time.Second)
	log.Print(timeout.Err())
}
