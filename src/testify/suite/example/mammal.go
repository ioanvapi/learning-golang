package animals

import (
	"fmt"
)

type Mammal interface {
	Move()
}
type Human struct {
	Name string
}

func (this *Human) Move() {
	fmt.Println(this.Name + " is walking...")
}
