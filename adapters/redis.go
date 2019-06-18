
package adapters

import "time"

// CacheAdapter - Adapter to talk to cache
type CacheAdapter interface {
	Get(key string) (string, error)
	Set(key string, value string, duration time.Duration) (string, error)
	Ping() error
}
