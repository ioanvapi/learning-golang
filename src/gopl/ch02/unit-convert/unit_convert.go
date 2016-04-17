package unit_convert

import "fmt"

type Feet float64
type Meter float64
type Pound float64
type Kilogram float64

func (f Feet) String() string     { return fmt.Sprintf("%g feet", f) }
func (m Meter) String() string    { return fmt.Sprintf("%g meter", m) }
func (p Pound) String() string    { return fmt.Sprintf("%g pound", p) }
func (k Kilogram) String() string { return fmt.Sprintf("%g kilogram", k) }

func FToM(f Feet) Meter { return Meter(f * 0.3048) }

func MToF(m Meter) Feet { return Feet(m / 0.3048) }

func PToK(p Pound) Kilogram { return Kilogram(p * 0.453592) }

func KToP(k Kilogram) Pound { return Pound(k / 0.453592) }
