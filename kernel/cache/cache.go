package cache

import (
	"gin.go.tpl/kernel/cfg"
	"github.com/gomodule/redigo/redis"
	"sync"
)

var (
	api  *Cache
	once sync.Once
)

func NewCacheApi(cfg *cfg.RedisConf) *Cache {
	once.Do(func() {
		api = &Cache{cfg: cfg}
	})
	return api
}

type Cache struct {
	cfg  *cfg.RedisConf
	pool *redis.Pool
}

func (c *Cache) initPool() {
	if c.pool == nil {
		c.pool = &redis.Pool{
			MaxIdle:   c.cfg.MaxIdle,
			MaxActive: c.cfg.MaxActive,
			Dial: func() (redis.Conn, error) {
				conn, err := redis.Dial(
					"tcp", c.cfg.Host+":"+c.cfg.Port,
					redis.DialPassword(c.cfg.Auth),
				)
				if err != nil {
					return nil, err
				}
				if c.cfg.DB != "" {
					_, _ = conn.Do("select", c.cfg.DB)
				}
				return conn, nil
			},
		}
	}
}

func (c *Cache) GetClient() redis.Conn {
	if c.pool == nil {
		c.initPool()
	}
	return c.pool.Get()
}
