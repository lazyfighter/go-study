package ratelimit_test

import (
	"context"
	rl "go-study/ratelimit"
	"go.uber.org/ratelimit"
	"golang.org/x/time/rate"
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

func TestGoLangLimit(t *testing.T) {
	limit := rate.Limit(10)
	limiter := rate.NewLimiter(limit, 100)
	for i := 1; i < 10; i++ {
		go func() {
			for {
				if limiter.Allow() {
					log.Println("allow")
				}
			}
		}()
	}
	time.Sleep(time.Minute)
}

func TestGolangWait(t *testing.T) {
	log.SetFlags(log.Lmicroseconds)
	limit := rate.Limit(10)
	count := 1
	limiter := rate.NewLimiter(limit, 1)
	for {
		timeout, _ := context.WithTimeout(context.Background(), time.Duration(100000*time.Millisecond))
		if limiter.Wait(timeout) == nil {
			count += 1
			log.Println(count)
		} else {
			log.Println("time out")
		}
	}

	time.Sleep(3 * time.Second)
}
