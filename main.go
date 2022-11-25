package main

import (
	"fmt"
	"time"
)

type Cacher interface {
	Get(string) (string, bool)
	Set(string, string)
}

// Тут надо дописать
type Cache struct {
	CacherTTL time.Duration
}

// Конструктор, тут надо дописать
func NewCache(ttl time.Duration) Cacher {
	m = make(map[string]data)
	return Cache{CacherTTL: ttl}
}

var m map[string]data

func (c Cache) Set(k1 string, v1 string) {
	m[k1] = data{value: v1, timeTTL: time.Now().Add(c.CacherTTL)}
}

func (c Cache) Get(k1 string) (string, bool) {
	if !time.Now().After(m[k1].timeTTL) {
		return m[k1].value, true
	}

	return m[k1].value, false
}

type data struct {
	value   string
	timeTTL time.Time
}

func main() {
	newCache := NewCache(5 * time.Second)
	// пустое значение
	v, ok := newCache.Get("test")
	fmt.Println(v)
	fmt.Println(ok)
	fmt.Println(`+++++`)
	newCache.Set("test", "test1")
	// не пустое значение
	time.Sleep(1 * time.Second)
	v, ok = newCache.Get("test")
	fmt.Println(v)
	fmt.Println(ok)
	fmt.Println(`+++++`)
	// не пустое значение через время
	time.Sleep(5 * time.Second)
	v, ok = newCache.Get("test")
	fmt.Println(v)
	fmt.Println(ok)
	fmt.Println(`+++++`)

}
