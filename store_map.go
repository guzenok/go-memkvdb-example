package memkvdb

import (
	"sync"
)

type MapStore struct {
	mutex sync.Mutex
	mem   map[DBKey][]byte
}

func CreateMapStore() MemStore {
	return &MapStore{
		mutex: sync.Mutex{},
		mem:   make(map[DBKey][]byte),
	}
}

func (s *MapStore) Set(key DBKey, val []byte) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.mem[key] = val

	return nil
}

func (s *MapStore) Get(key DBKey) ([]byte, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	val, ok := s.mem[key]
	if !ok {
		return nil, ErrNotFound
	} else {
		delete(s.mem, key)
		return val, nil
	}
}

func (s *MapStore) Del(key DBKey) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	delete(s.mem, key)
}
