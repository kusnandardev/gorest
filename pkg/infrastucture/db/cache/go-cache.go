package cache

import (
	"github.com/patrickmn/go-cache"
	"time"
)

type Cache struct {
	c *cache.Cache
}

func NewCache() *Cache {
	return &Cache{}
}

func (c Cache) SetCache() {
	c.c = cache.New(5*time.Minute, 10*time.Minute)
}

func (c *Cache) GetCache() *cache.Cache {
	return c.c
}
