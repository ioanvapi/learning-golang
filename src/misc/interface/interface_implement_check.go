package main

type Jedi interface {
	HasForce() bool
}

type Knight struct {
}

func (k *Knight) HasForce() bool {
	return true
}

var _ Jedi = (*Knight)(nil)

func main() {
}
