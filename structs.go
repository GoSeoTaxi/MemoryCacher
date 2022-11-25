package MemoryCacher

import (
	"sync"
	"time"
)

type data struct {
	value   interface{}
	timeTTL time.Time
}

type Cache struct {
	cacherTTL time.Duration
	storage   *sync.Map
}
