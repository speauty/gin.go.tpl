package cache

import (
	"gin.go.tpl/kernel/cfg"
	"github.com/go-redis/redis/v8"
	"net"
	"sync"
)

var (
	api  *Cache
	once sync.Once
)

func NewCacheApi(cfg *cfg.RedisConf) *Cache {
	once.Do(func() {
		api = &Cache{cfg: cfg}
		api.initClient()
	})
	return api
}

type Cache struct {
	cfg   *cfg.RedisConf
	redis *redis.Client
}

func (c *Cache) NewClient(currentCfg *cfg.RedisConf) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     net.JoinHostPort(currentCfg.Host, currentCfg.Port),
		Password: currentCfg.Auth,
		DB:       currentCfg.DB,
	})
}

func (c *Cache) GetClient() *redis.Client {
	if c.redis == nil {
		c.initClient()
	}
	return c.redis
}

func (c *Cache) initClient() {
	if c.redis == nil {
		c.redis = c.NewClient(c.cfg)
	}
}
