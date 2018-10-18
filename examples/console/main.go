package main

import (
	"time"

	db "github.com/guzenok/go-memkvdb-example"
)

func main() {

	memcache, err := db.New(30*time.Second, db.CreateMapStore())
	if err != nil {
		panic(err)
	}

	key := []byte("index")
	val := []byte("stored data")

	err = memcache.Set(key, val)
	if err != nil {
		panic(err)
	}

	val, err = memcache.Get(key)
	switch err {
	case nil:
		// OK
		break
	case db.ErrNotFound:
		// key not found or expired
		break
	default:
		panic(err)
	}
}
