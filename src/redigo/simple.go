package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	rdsConn, gErr := redis.Dial("tcp", "0.0.0.0:6379")
	if gErr != nil {
		fmt.Println(gErr)
		return
	}
	defer rdsConn.Close()

	if _, dErr := rdsConn.Do("SET", "a", "apple"); gErr != nil {
		fmt.Println(dErr)
		return
	}

	if reply, rErr := rdsConn.Do("GET", "a"); rErr != nil {
		fmt.Println(rErr)
		return
	} else {
		if replyBytes, ok := reply.([]byte); ok {
			fmt.Println(string(replyBytes))
		} else {
			fmt.Println("Err: get value by string key")
		}
	}
}
