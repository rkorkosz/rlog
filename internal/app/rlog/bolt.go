package rlog

import (
	"encoding/json"

	bolt "go.etcd.io/bbolt"
)

type BoltStorage struct {
	path   string
	bucket []byte
}

func NewBoltStorage(path, bucket string) (*BoltStorage, error) {
	db, err := bolt.Open(path, 0600, nil)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucket))
		return err
	})
	return &BoltStorage{path: path, bucket: []byte(bucket)}, err
}

func (b *BoltStorage) Get(slug string) (*Entry, error) {
	db, err := bolt.Open(b.path, 0600, nil)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	e := Entry{}
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(b.bucket)
		v := b.Get([]byte(slug))
		return json.Unmarshal(v, &e)
	})
	return &e, err
}

func (b *BoltStorage) Put(e Entry) error {
	db, err := bolt.Open(b.path, 0600, nil)
	if err != nil {
		return err
	}
	defer db.Close()
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(b.bucket)
		jn, err := json.Marshal(&e)
		if err != nil {
			return err
		}
		return b.Put([]byte(e.Slug), jn)
	})
}

func (b *BoltStorage) Delete(slug string) error {
	db, err := bolt.Open(b.path, 0600, nil)
	if err != nil {
		return err
	}
	defer db.Close()
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(b.bucket)
		return b.Delete([]byte(slug))
	})
}

func (b *BoltStorage) List() ([]Entry, error) {
	db, err := bolt.Open(b.path, 0600, nil)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	entries := []Entry{}
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(b.bucket)
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			e := Entry{}
			err := json.Unmarshal(v, &e)
			if err != nil {
				return err
			}
			entries = append(entries, e)
		}
		return nil
	})
	return entries, err
}
