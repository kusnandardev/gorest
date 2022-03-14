package inmemory

import (
	"RestGo/pkg/infrastucture/db/inmemory"
	"errors"
	"github.com/patrickmn/go-cache"
	"time"
)

type CacheHandler struct {
	cache *cache.Cache
}

var cc *cache.Cache

func DefaultCacheHandler() *CacheHandler {
	cc = inmemory.GetCache()
	return &CacheHandler{cache: cc}
}

func (h *CacheHandler) Get(key string) (interface{}, error) {
	value, ok := h.cache.Get(key)
	if !ok {
		return "", errors.New("no value")
	}
	return value, nil
}

func (h *CacheHandler) Set(key string, value interface{}) error {
	h.cache.Set(key, value, (5 * time.Minute))
	return nil
}

func (h *CacheHandler) Delete(key string) error {
	h.cache.Delete(key)
	return nil
}
