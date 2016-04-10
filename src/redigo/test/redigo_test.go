package redis

import (
	"testing"
	"time"

	"github.com/garyburd/redigo/redis"
)

func BenchmarkNoPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		func() {
			rdsConn, gErr := redis.Dial("tcp", "0.0.0.0:6379")
			if gErr != nil {
				b.Error(gErr)
				b.FailNow()
			}
			defer rdsConn.Close()

			if _, dErr := rdsConn.Do("SET", "a", "apple"); dErr != nil {
				b.Error(dErr)
				b.FailNow()
			}

			if reply, rErr := rdsConn.Do("GET", "a"); rErr != nil {
				b.Error(rErr)
				b.FailNow()
			} else {
				if _, ok := reply.([]byte); ok {
					//b.Log(string(replyBytes))
				} else {
					b.Error("Err: get value by string key")
				}
			}
		}()
	}
}

func BenchmarkWithPool(b *testing.B) {
	rdsPool := &redis.Pool{
		MaxIdle:     100,
		IdleTimeout: time.Second * 300,
		Dial: func() (redis.Conn, error) {
			conn, cErr := redis.Dial("tcp", "0.0.0.0:6379")
			if cErr != nil {
				return nil, cErr
			}
			return conn, nil
		},
	}

	for i := 0; i < b.N; i++ {
		func() {
			rdsConn := rdsPool.Get()
			defer rdsConn.Close()

			if _, dErr := rdsConn.Do("SET", "a", "apple"); dErr != nil {
				b.Error(dErr)
				b.FailNow()
			}

			if reply, rErr := rdsConn.Do("GET", "a"); rErr != nil {
				b.Error(rErr)
				b.FailNow()
			} else {
				if _, ok := reply.([]byte); ok {
					//b.Log(string(replyBytes))
				} else {
					b.Error("Err: get value by string key")
				}
			}
		}()
	}
}
