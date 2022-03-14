package container

import (
	"github.com/patrickmn/go-cache"
	"github.com/sarulabs/di"
	"time"
)

func MainCacheContainer(ctn di.Container) (interface{}, error) {
	cc := cache.New(5*time.Minute, 10*time.Minute)
	return cc, nil
}
