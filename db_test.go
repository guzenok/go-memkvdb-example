package memkvdb

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDb(t *testing.T) {
	assert := assert.New(t)

	db, err := New(30 * time.Second)
	assert.NotNil(db)
	assert.Nil(err)

	err = db.Set(nil, []byte{0xff})
	assert.Equal(ErrEmptyKey, err)

	err = db.Set([]byte{}, []byte{0x00})
	assert.Equal(ErrEmptyKey, err)

	// test data
	data := []struct {
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

	// set values
	for _, d := range data {
		err = db.Set(d.key, d.val)
		assert.Nil(err)
	}

	var val []byte

	// get not existing
	for _, d := range data[len(data)-1:] {
		val, err = db.Get(d.val)
		assert.Equal(err, ErrNotFound)
		assert.Nil(val)
		break
	}

	// get existing
	for _, d := range data {
		val, err = db.Get(d.key)
		assert.Nil(err)
		assert.Equal(d.val, val)
	}

	// get deleted
	for _, d := range data {
		val, err = db.Get(d.key)
		assert.Equal(err, ErrNotFound)
		assert.Nil(val)
	}

}
