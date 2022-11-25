// Memory cacher key-value
// [string]interface{}
package MemoryCacher

import (
	"sync"
	"time"
)

func NewCache(ttl time.Duration) Cacher {
	sm := sync.Map{}
	go func(m *sync.Map) {
		for {
			time.Sleep(ttl + (1 * time.Second))
			m.Range(func(key interface{}, value interface{}) bool {
				if time.Now().After(value.(data).timeTTL) {
					m.Delete(key)
				}
				return true
			})
			//	fmt.Println(`clean complete`)
		}
	}(&sm)
	return Cache{cacherTTL: ttl, storage: &sm}
}
