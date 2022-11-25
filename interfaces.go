package MemoryCacher

import "time"

type Cacher interface {
	Get(string) (interface{}, bool)
	Set(string, interface{})
}

func (c Cache) Set(key string, value interface{}) {
	c.storage.Store(key, data{value: value, timeTTL: time.Now().Add(c.cacherTTL)})
}

func (c Cache) Get(k1 string) (interface{}, bool) {
	item, ok := c.storage.Load(k1)
	if ok && !time.Now().After(item.(data).timeTTL) {
		return item.(data).value, true
	}
	return nil, false
}
