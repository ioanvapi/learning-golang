// http://www.infoq.com/cn/articles/go-interface-talk
package main

import (
	"fmt"
)

type Err struct{}

func (_ *Err) Error() string {
	return "To err is human"
}

func NoErr(ok bool) error {
	if !ok {
		return &Err{}
	}
	return nil
}

func main() {
	fmt.Println(NoErr(true))
	fmt.Println(NoErr(false))
}
