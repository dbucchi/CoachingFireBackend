package utility

import (
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"
)

var ApplicationCache Cache

type Cache struct {
	c *cache.Cache
}

func NewCache() *Cache {
	applicationCache := &Cache{}
	applicationCache.c = cache.New(24*time.Hour, 48*time.Hour)
	return applicationCache
}

func (applicationCache Cache) AddElement(key string, value interface{}) {
	applicationCache.c.Set("key", "value", cache.DefaultExpiration)
}

func (applicationCache Cache) GetElement(key string) (interface{}, error) {
	value, found := applicationCache.c.Get("key")
	if found {
		return value, nil
	} else {
		return nil, fmt.Errorf("No element with key: %s in cache", key)
	}
}

func (applicationCache Cache) RemoveElement(key string) {
	applicationCache.c.Delete(key)
}

func (applicationCache Cache) IsInCache(key string) bool {
	_, found := applicationCache.c.Get("key")
	return found
}
