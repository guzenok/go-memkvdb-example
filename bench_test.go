package memkvdb

import (
	"math/rand"
	"testing"
	"time"
)

const (
	KEY_LEN   = 256
	VAL_LEN   = 2048
	READ_RATE = 10
)

func BenchmarkDB(b *testing.B) {

	key := make([]byte, KEY_LEN)
	val := make([]byte, VAL_LEN)

	_, err := rand.Read(key[:])
	if err != nil {
		b.Fatal(err)
	}
	_, err = rand.Read(val[:])
	if err != nil {
		b.Fatal(err)
	}

	for name, store := range map[string]MemStore{
		"Map":     CreateMapStore(),
		"SyncMap": CreateSyncMapStore()} {
		b.Run(name, func(b *testing.B) {

			expiration := 30 * time.Second
			db, err := New(expiration, store)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				key[i%KEY_LEN] = byte(rand.Int())

				err = db.Set(key, val)
				if err != nil {
					b.Fatal(err)
				}

				if i%READ_RATE != 0 {
					continue
				}

				_, err = db.Get(key)
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}
