package rlog

import (
	"errors"
	"sync"
)

type Cache struct {
	mu   sync.RWMutex
	db   map[string]Entry
	path string
}

func NewCache(path string) *Cache {
	return &Cache{
		mu:   sync.RWMutex{},
		db:   make(map[string]Entry),
		path: path,
	}
}

func (c *Cache) Get(slug string) (*Entry, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	e, ok := c.db[slug]
	if !ok {
		return nil, errors.New("not found")
	}
	return &e, nil
}

func (c *Cache) Put(e Entry) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.db[e.Slug] = e
	return nil
}

func (c *Cache) Delete(slug string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.db, slug)
	return nil
}

func (c *Cache) List() ([]Entry, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	out := []Entry{}
	for _, v := range c.db {
		out = append(out, v)
	}
	return out, nil
}
