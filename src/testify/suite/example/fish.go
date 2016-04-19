package animals

import (
	"fmt"
)

type ClownFish struct {
	Name string
}

func (this *ClownFish) Swim() {
	fmt.Println(this.Name + " is swimming...")
}

type Shark struct {
	Name string
}

func (this *Shark) Swim() {
	fmt.Println(this.Name + " is swimming...")
}
