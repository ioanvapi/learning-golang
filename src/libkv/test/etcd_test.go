package etcd_test

import (
	"testing"
	"time"

	"github.com/docker/libkv"
	"github.com/docker/libkv/store"
	"github.com/docker/libkv/store/etcd"
)

func init() {
	etcd.Register()
}

func BenchmarkSetGetDelete(b *testing.B) {
	b.N = 1000
	for i := 0; i < b.N; i++ {
		func() {
			client := "localhost:2379"

			// Initialize a new store with etcd
			kv, err := libkv.NewStore(
				store.ETCD, // or "etcd"
				[]string{client},
				&store.Config{
					ConnectionTimeout: 10 * time.Second,
				},
			)
			if err != nil {
				b.Fatal("Cannot create store consul")
			}

			key := "foo"
			err = kv.Put(key, []byte("bar"), nil)
			if err != nil {
				b.Errorf("Error trying to put value at key: %v", key)
			}

			pair, err := kv.Get(key)
			if err != nil {
				b.Errorf("Error trying accessing value at key: %v", key)
			}

			err = kv.Delete(key)
			if err != nil {
				b.Errorf("Error trying to delete key %v", key)
			}
			_ = pair
			// b.Errorf("value: %v", string(pair.Value))
		}()
	}
}
