package helpers

import (
	"time"

	cache "github.com/patrickmn/go-cache"
)

// CacheInstance is an instance of cache.
var CacheInstance = cache.New(20*time.Second, 72*time.Hour)
