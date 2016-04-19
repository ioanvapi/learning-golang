package animals

import (
	"fmt"
)

type Bird struct {
	Name string
}

func (this *Bird) Fly() {
	fmt.Println(this.Name + " is flying...")
}
