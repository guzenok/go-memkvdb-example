package memkvdb

import (
	"errors"
	"sync"
	"time"
)

var (
	ErrNotFound = errors.New("key not found")
)

type DB struct {
	expiration time.Duration
	mutex      sync.RWMutex
	memstore   map[DBKey][]byte
}

func New(expiration time.Duration) (*DB, error) {
	db := &DB{
		expiration: expiration,
		mutex:      sync.RWMutex{},
		memstore:   make(map[DBKey][]byte),
	}, nil

	return db
}

func (db *DB) Set(key, val []byte) error {
	hash, err := hash(key)
	if err != nil {
		return err
	}

	db.mutex.Lock()
	defer db.mutex.Unlock()

	db.memstore[hash] = val
	return nil
}

func (db *DB) Get(key []byte) ([]byte, error) {
	hash, err := hash(key)
	if err != nil {
		return nil, err
	}

	db.mutex.RLock()
	defer db.mutex.RUnlock()

	val, ok := db.memstore[hash]
	if !ok {
		return nil, ErrNotFound
	} else {
		delete(db.memstore, hash)
		return val, nil
	}
}
