package memkvdb

import (
	"sync"
	"time"
)

type DB struct {
	expiration time.Duration
	mutex      sync.RWMutex
	memstore   map[DBKey][]byte
}

func New(expiration time.Duration) (*DB, error) {
	return &DB{
		expiration: expiration,
		mutex:      sync.RWMutex{},
		memstore:   make(map[DBKey][]byte),
	}, nil
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

	return db.memstore[hash], nil
}
