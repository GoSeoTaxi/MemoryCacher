//Memory cacher key-value
// [string]interface{}

package github.com/GoSeoTaxi/MemoryCacher

import (
	"fmt"
	"sync"
	"time"
)

type Cacher interface {
	Get(string) (interface{}, bool)
	Set(string, interface{})
}

type Cache struct {
	cacherTTL time.Duration
	storage   *sync.Map
}

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
			fmt.Println(`clean complete`)
		}
	}(&sm)
	return Cache{cacherTTL: ttl, storage: &sm}
}

func (c Cache) Set(key string, value interface{}) {
	c.storage.Store(key, data{value: value, timeTTL: time.Now().Add(c.cacherTTL)})
}

func (c Cache) Get(k1 string) (interface{}, bool) {
	item, ok := c.storage.Load(k1)
	if ok && !time.Now().After(item.(data).timeTTL) {
		return item.(data).value, true
	}
	return "", false
}

type data struct {
	value   interface{}
	timeTTL time.Time
}
