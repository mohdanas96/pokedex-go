package pokecache

import "time"

type CacheEntry struct {
	createdAt time.Time
	data      []byte
}
