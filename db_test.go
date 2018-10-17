package memkvdb

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	testData = []struct {
		key []byte
		val []byte
	}{
		{
			key: []byte{0x93, 0xfe, 0xd, 0xa0, 0x89, 0xcd, 0xf6, 0x34, 0x90, 0x4f, 0xd5, 0x9f, 0x71},
			val: nil,
		},
		{
			key: []byte{0x93, 0xb8, 0x85, 0xad, 0xfe, 0xd, 0xa0, 0x89, 0xcd, 0x90, 0x4f, 0xd5, 0x9f, 0x71},
			val: []byte{},
		},
		{
			key: []byte{0x00},
			val: []byte{0x93, 0xfe, 0xd, 0xa0, 0x89, 0xcd, 0x93, 0xb8, 0x85, 0xad, 0xfe, 0xd, 0xa0, 0x89, 0xf6, 0x34, 0x90, 0x4f, 0xd5, 0x9f, 0x71},
		},
	}
)

func TestEmptyKey(t *testing.T) {

	for name, store := range map[string]MemStore{
		"Map":     CreateMapStore(),
		"SyncMap": CreateSyncMapStore()} {
		t.Run(name, func(t *testing.T) {

			expiration := 30 * time.Second

			assert := assert.New(t)
			db := initDB(assert, expiration, store)

			err := db.Set(nil, []byte{0xff})
			assert.Equal(ErrEmptyKey, err)

			err = db.Set([]byte{}, []byte{0x00})
			assert.Equal(ErrEmptyKey, err)

			_, err = db.Get(nil)
			assert.Equal(ErrEmptyKey, err)

			_, err = db.Get([]byte{})
			assert.Equal(ErrEmptyKey, err)
		})
	}
}

func TestRW(t *testing.T) {

	for name, store := range map[string]MemStore{
		"Map":     CreateMapStore(),
		"SyncMap": CreateSyncMapStore()} {
		t.Run(name, func(t *testing.T) {

			expiration := 30 * time.Second

			assert := assert.New(t)
			db := initDB(assert, expiration, store)

			var val []byte
			var err error

			// write
			for _, d := range testData {
				err = db.Set(d.key, d.val)
				assert.Nil(err)
			}

			// read not existing
			for _, d := range testData[len(testData)-1:] {
				val, err = db.Get(d.val)
				assert.Equal(ErrNotFound, err)
				assert.Nil(val)
				break
			}

			// read existing
			for _, d := range testData {
				val, err = db.Get(d.key)
				assert.Nil(err)
				assert.Equal(d.val, val)
			}

			// read deleted
			for _, d := range testData {
				val, err = db.Get(d.key)
				assert.Equal(ErrNotFound, err)
				assert.Nil(val)
			}
		})
	}
}

func TestTTL(t *testing.T) {
	for name, store := range map[string]MemStore{
		"Map":     CreateMapStore(),
		"SyncMap": CreateSyncMapStore()} {
		t.Run(name, func(t *testing.T) {

			expiration := 200 * time.Millisecond
			wait_time := 210 * time.Millisecond

			assert := assert.New(t)
			db := initDB(assert, expiration, store)

			var val []byte
			var err error

			for _, d := range testData {
				err = db.Set(d.key, d.val)
				assert.Nil(err)
			}

			for _, d := range testData[0:1] {
				val, err = db.Get(d.key)
				assert.Nil(err)
				assert.Equal(d.val, val)
			}

			<-time.After(wait_time)

			for _, d := range testData[1:] {
				val, err = db.Get(d.key)
				assert.Equal(ErrNotFound, err)
				assert.Nil(val)
			}
		})
	}
}

func initDB(assert *assert.Assertions, expiration time.Duration, store MemStore) *DB {
	db, err := New(expiration, store)
	assert.NotNil(db)
	assert.Nil(err)
	return db
}
