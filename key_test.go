package memkvdb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	assert := assert.New(t)

	res, err := hash(nil)
	assert.Equal(DBKey(""), res)
	assert.Equal(ErrEmptyKey, err)

	res, err = hash([]byte{})
	assert.Equal(DBKey(""), res)
	assert.Equal(ErrEmptyKey, err)

	res, err = hash([]byte{0x00})
	assert.Equal(DBKey("AA=="), res)
	assert.Nil(err)

	res, err = hash([]byte{0xff, 0x00, 0x00, 0x00, 0xf0, 0x1e, 0x0a, 0x99, 0xd4, 0x1d, 0x8c, 0xd9, 0x8f, 0x0, 0xb2, 0x4, 0xe9, 0x80, 0x9, 0x98, 0xec, 0xf8, 0x42, 0x7e})
	assert.Equal(DBKey("/wAAAPAeCpnUHYzZjwCyBOmACZjs+EJ+"), res)
	assert.Nil(err)

}
