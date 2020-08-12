package cache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var Cache *cache.Cache

func init() {
	Cache = cache.New(10*time.Minute, 1*time.Minute)
}
