package db

import (
	"time"

	"github.com/patrickmn/go-cache"
)

// Create a cache with a default expiration time of 30 minutes, and which
// purges expired items every 20 minutes
var Database = cache.New(30*time.Minute, 20*time.Minute)
