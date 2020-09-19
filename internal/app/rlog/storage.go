package rlog

import (
	"encoding/json"
	"errors"
	"os"
	"sync"
)

type EntryStorage interface {
	Get(slug string) (*Entry, error)
	Put(e Entry) error
	Delete(slug string) error
	List() ([]Entry, error)
}

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

func (c *Cache) Dump() error {
	c.mu.RLock()
	defer c.mu.RUnlock()
	file, err := os.OpenFile(c.path, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0755)
	if err != nil {
		return nil
	}
	defer file.Close()
	return json.NewEncoder(file).Encode(&c.db)
}

func (c *Cache) Load() error {
	c.mu.Lock()
	defer c.mu.Unlock()
	file, err := os.Open(c.path)
	if err != nil {
		return nil
	}
	defer file.Close()
	return json.NewDecoder(file).Decode(&c.db)
}
