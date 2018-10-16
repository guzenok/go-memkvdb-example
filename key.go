package memkvdb

import (
	"crypto/md5"
	"errors"
)

var (
	ErrEmptyKey = errors.New("key is empty")
)

const keyLen = 16

type DBKey [keyLen]byte

func hash(key []byte) (res DBKey, err error) {
	if key == nil || len(key) < 1 {
		err = ErrEmptyKey
		return
	}

	hasher := md5.New()
	_, err = hasher.Write(key)
	if err != nil {
		return
	}
	h := hasher.Sum(nil)
	copy(res[:], h[0:keyLen])
	return
}
