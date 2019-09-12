package ratelimit

import (
	"log"
	"sync"
	"time"
)

type Limiter interface {
	Take() bool
}

type limiter struct {
	sync.Mutex
	bucket int64
}

func (l *limiter) Take() bool {
	l.Lock()
	defer l.Unlock()
	if l.bucket > 0 {
		l.bucket -= 1
		return true
	}
	return false
}

func NewTokenBucket(rate int64) Limiter {
	perRequest := time.Second / time.Duration(rate)
	log.Println(perRequest)
	ticker := time.Tick(perRequest)
	l := &limiter{
		bucket: 0,
	}

	go func() {
		for {
			<-ticker
			l.Lock()
			if l.bucket < rate {
				l.bucket = l.bucket + 1
			}
			l.Unlock()
		}
	}()
	return l
}
