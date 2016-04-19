package animals

type Animal interface {
	Move() // any type that implements an area method is considered an Animal
}

type MultiAnimal struct {
	animals []Animal // animals field is a slice of interfaces
}

func (m *MultiAnimal) Move() {
	for _, animal := range m.animals { // iterate through animals
		animal.Move() // execute polymorphic Move method for this Animal
	}
}

type ClownFishAdapter struct {
	*ClownFish
}

func (this *ClownFishAdapter) Move() {
	this.Swim()
}

type SharkAdapter struct {
	*Shark
}

func (this *SharkAdapter) Move() {
	this.Swim()
}

type BirdAdapter struct {
	*Bird
}

func (this *BirdAdapter) Move() {
	this.Fly()
}
