package limiter

import (
	"golang.org/x/time/rate"
	"sync"
	"time"
)

type Limiters struct {
	limiters *sync.Map
}

type Limiter struct {
	limiter   *rate.Limiter
	lastGetAt time.Time //上一次获取token的时间
	key       string
}

var GlobalLimiters = &Limiters{
	limiters: &sync.Map{},
}

var once = sync.Once{}

func NewLimiter(rateInterval rate.Limit, tokenGenerated int, key string) *Limiter {
	once.Do(func() {
		go GlobalLimiters.clearLimiter()
	})
	keyLimiter := GlobalLimiters.getLimiter(rateInterval, tokenGenerated, key)
	return keyLimiter

}

func (l *Limiter) Allow() bool {
	l.lastGetAt = time.Now()
	return l.limiter.Allow()

}

func (ls *Limiters) getLimiter(rateInterval rate.Limit, tokenGenerated int, key string) *Limiter {
	limiter, ok := ls.limiters.Load(key)
	if ok {
		return limiter.(*Limiter)
	}

	l := &Limiter{
		limiter:   rate.NewLimiter(rateInterval, tokenGenerated),
		lastGetAt: time.Now(),
		key:       key,
	}

	ls.limiters.Store(key, l)
	return l
}

// 清除过期的限流器(超过一分钟)
func (ls *Limiters) clearLimiter() {
	for {
		time.Sleep(1 * time.Minute)
		ls.limiters.Range(func(key, value interface{}) bool {
			if time.Now().Unix()-value.(*Limiter).lastGetAt.Unix() > 60 {
				ls.limiters.Delete(key)
			}
			return true
		})

	}

}
