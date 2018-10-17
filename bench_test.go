package memkvdb

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

const (
	KEY_LEN   = 256
	VAL_LEN   = 2048
	READ_RATE = 10
)

func BenchmarkDbSync(b *testing.B) {
	val := make([]byte, VAL_LEN)
	rand.Read(val[:])

	key := make([]byte, KEY_LEN)
	rand.Read(key[:])

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

func BenchmarkDbAsync(b *testing.B) {
	val := make([]byte, VAL_LEN)
	rand.Read(val[:])

	for name, store := range map[string]MemStore{
		"Map":     CreateMapStore(),
		"SyncMap": CreateSyncMapStore()} {
		b.Run(name, func(b *testing.B) {

			expiration := 30 * time.Second
			db, err := New(expiration, store)
			wg := &sync.WaitGroup{}

			b.ResetTimer()
			for i := 0; i < b.N/2/READ_RATE; i++ {
				wg.Add(1)
				go func(i int) {
					defer wg.Done()

					key := make([]byte, KEY_LEN)
					rand.Read(key[:])

					for j := 0; j < 2*READ_RATE; j++ {
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
				}(i)
			}
			wg.Wait()
		})
	}
}
