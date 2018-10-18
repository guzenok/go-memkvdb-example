package memkvdb

import (
	"encoding/base64"
	"errors"
)

var (
	ErrEmptyKey = errors.New("key is empty")
)

type DBKey string

func hash(key []byte) (res DBKey, err error) {
	if len(key) < 1 {
		err = ErrEmptyKey
		return
	}

	res = DBKey(base64.StdEncoding.EncodeToString(key))
	return
}
