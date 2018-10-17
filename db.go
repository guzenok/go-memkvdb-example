package memkvdb

import (
	"context"
	"errors"
	"time"
)

var (
	ErrNotFound = errors.New("key not found")
)

type MemStore interface {
	Set(key DBKey, val []byte) error
	Get(key DBKey) ([]byte, error)
	Del(key DBKey)
}

type DB struct {
	expiration time.Duration
	store      MemStore
}

func New(expiration time.Duration, store MemStore) (*DB, error) {
	db := &DB{
		expiration: expiration,
		store:      store,
	}
	return db, nil
}

func NewDefault(expiration time.Duration) (*DB, error) {
	var store MemStore
	store = CreateMapStore()

	return New(expiration, store)
}

func (db *DB) Set(key, val []byte) error {
	hash, err := hash(key)
	if err != nil {
		return err
	}

	err = db.store.Set(hash, val)
	if err != nil {
		return err
	}

	// TTL
	ctx, _ := context.WithTimeout(context.Background(), db.expiration)
	go func() {
		<-ctx.Done()
		db.store.Del(hash)
	}()

	return nil
}

func (db *DB) Get(key []byte) ([]byte, error) {
	hash, err := hash(key)
	if err != nil {
		return nil, err
	}

	return db.store.Get(hash)
}
