package localcache

import (
	"github.com/coocood/freecache"
)

const (
	CACHE_SIZE = 1024 * 1024 * 1024
)

func InitLocalCache() {
	freecache.NewCache(CACHE_SIZE)
}