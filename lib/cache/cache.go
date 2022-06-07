package cache

import (
	"gin.go.tpl/lib/config"
	"github.com/gomodule/redigo/redis"
	"sync"
)

var (
	CacheApi  *Cache
	CacheOnce sync.Once
)

type Cache struct {
	config config.RedisConf
	Pool   *redis.Pool
}

func NewCacheApi(config config.RedisConf) *Cache {
	CacheOnce.Do(func() {
		CacheApi = &Cache{config: config}
	})
	return CacheApi
}

func (c *Cache) initPool() {
	if c.Pool == nil {
		c.Pool = &redis.Pool{
			MaxIdle:   c.config.MaxIdle,
			MaxActive: c.config.MaxActive,
			Dial: func() (redis.Conn, error) {
				conn, err := redis.Dial(
					"tcp", c.config.Host+":"+c.config.Port,
					redis.DialPassword(c.config.Auth),
				)
				if err != nil {
					return nil, err
				}
				if c.config.DB != "" {
					_, _ = conn.Do("select", c.config.DB)
				}
				return conn, nil
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
