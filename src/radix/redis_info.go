package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/mediocregopher/radix.v2/redis"
)

func errHndlr(err error) {
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}

// BuildMapFromInfoString will take the string from a Redis info call and
// return a map[string]string
func BuildMapFromInfoString(input string) map[string]string {
	imap := make(map[string]string)
	lines := strings.Split(input, "\r\n")
	for _, line := range lines {
		if len(line) > 0 {
			if strings.Contains(line, "#") {
				imap["section"] = strings.Split(line, "#")[1]
			} else {
				splits := strings.Split(line, ":")
				key := splits[0]
				val := splits[1]
				imap[key] = val
			}
		}
	}
	return imap
}

func main() {
	c, err := redis.DialTimeout("tcp", "127.0.0.1:6379", 10*time.Second)
	errHndlr(err)
	defer c.Close()

	r := c.Cmd("select", 0)
	errHndlr(r.Err)

	r = c.Cmd("info", "replication")
	errHndlr(r.Err)

	ret, _ := r.Str()
	fmt.Println(ret)
	m := BuildMapFromInfoString(ret)
	fmt.Println(m)

	r = c.Cmd("info", "server")
	errHndlr(r.Err)

	ret, _ = r.Str()
	fmt.Println(ret)
	m = BuildMapFromInfoString(ret)
	fmt.Println(m)

}
