package memkvdb

import (
	"sync"
)

type SyncMapStore struct {
	mem *sync.Map
}

func CreateSyncMapStore() MemStore {
	return &SyncMapStore{
		mem: &sync.Map{},
	}
}

func (s *SyncMapStore) Set(key DBKey, val []byte) error {
	s.mem.Store(key, val)
	return nil
}

func (s *SyncMapStore) Get(key DBKey) ([]byte, error) {
	val, ok := s.mem.Load(key)
	if !ok {
		return nil, ErrNotFound
	} else {
		s.mem.Delete(key)
		return val.([]byte), nil
	}
}

func (s *SyncMapStore) Del(key DBKey) {
	s.mem.Delete(key)
}
