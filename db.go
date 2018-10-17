package memkvdb

import (
	"context"
	"errors"
	"sync"
	"time"
)

var (
	ErrNotFound = errors.New("key not found")
)

type DB struct {
	expiration time.Duration
	mutex      sync.Mutex
	memstore   map[DBKey][]byte
}

func New(expiration time.Duration) (*DB, error) {
	db := &DB{
		expiration: expiration,
		mutex:      sync.Mutex{},
		memstore:   make(map[DBKey][]byte),
	}

	return db, nil
}

func (db *DB) Set(key, val []byte) error {
	hash, err := hash(key)
	if err != nil {
		return err
	}

	db.mutex.Lock()
	defer db.mutex.Unlock()

	db.memstore[hash] = val

	// TTL
	ctx, _ := context.WithTimeout(context.Background(), db.expiration)
	go func() {
		<-ctx.Done()
		db.get(hash)
	}()

	return nil
}

func (db *DB) Get(key []byte) ([]byte, error) {
	hash, err := hash(key)
	if err != nil {
		return nil, err
	}

	return db.get(hash)
}

func (db *DB) get(hash DBKey) ([]byte, error) {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	val, ok := db.memstore[hash]
	if !ok {
		return nil, ErrNotFound
	} else {
		delete(db.memstore, hash)
		return val, nil
	}
}
