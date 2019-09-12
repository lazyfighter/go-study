package ratelimit_test

import (
	rl "go-study/src/ratelimit"
	"go.uber.org/ratelimit"
	"log"
	"testing"
	"time"
)

func TestUberRateLimit(t *testing.T) {
	limiter := ratelimit.New(60000)
	now := time.Now()
	for i := 0; i < 10; i++ {
		take := limiter.Take()
		log.Println(take.Sub(now))
		now = time.Now()
	}
}

func TestUberUnLimit(t *testing.T) {
	unlimited := ratelimit.NewUnlimited()
	for i := 0; i < 10; i++ {
		log.Println(unlimited.Take())
	}
}

func TestNewTokenBucket(t *testing.T) {
	limiter := rl.NewTokenBucket(100)
	log.Println(limiter)
	pre := time.Now()
	//time.Sleep(10 * time.Second)
	for {
		if limiter.Take() {
			log.Println(time.Now().Sub(pre))
			pre = time.Now()
		}
	}

}
