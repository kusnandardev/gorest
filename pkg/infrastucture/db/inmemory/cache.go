package inmemory

import (
	"github.com/patrickmn/go-cache"
	"time"
)

var cc *cache.Cache

func InitCache() {
	cc = cache.New(5*time.Minute, 10*time.Minute)
}

func GetCache() *cache.Cache {
	return cc
}
