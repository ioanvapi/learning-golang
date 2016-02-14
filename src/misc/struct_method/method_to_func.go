package main

type Data struct{}

func (Data) TestValue()    {}
func (*Data) TestPointer() {}

func main() {
	var p *Data = nil
	p.TestPointer()

	(*Data)(nil).TestPointer() // method value
	(*Data).TestPointer(nil)   // method expression
	// p.TestValue()
	// (Data)(nil).TestValue()
	// Data.TestValue(nil)
}
