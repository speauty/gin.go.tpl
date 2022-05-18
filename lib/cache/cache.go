package cache

import (
	"context"
	"github.com/gomodule/redigo/redis"
	"sync"
)

var (
	CacheAPI  *Cache
	CacheOnce sync.Once
)

type Cache struct {
	bg   *context.Context
	Pool *redis.Pool
}

func NewCacheAPI(bg context.Context) *Cache {
	CacheOnce.Do(func() {
		CacheAPI = &Cache{
			bg: &bg,
		}
	})
	return CacheAPI
}

func (c *Cache) initPool() {
	if c.Pool == nil {
		c.Pool = &redis.Pool{
			MaxIdle:   10, /*最大的空闲连接数*/
			MaxActive: 10, /*最大的激活连接数*/
			Dial: func() (redis.Conn, error) {
				//c, err := redis.Dial("tcp", c.Conf.Host+":"+c.Conf.Port, redis.DialPassword(c.Conf.Auth))
				//if err != nil {
				//	return nil, err
				//}
				return nil, nil
			},
		}
	}
}

func (c *Cache) GetClient() redis.Conn {
	if c.Pool == nil {
		c.initPool()
	}
	return c.Pool.Get()
}
