package weightconv

import "fmt"

type Pound float64
type Kilogramme float64

func (p Pound) String() string      { return fmt.Sprintf("%glb", p) }
func (k Kilogramme) String() string { return fmt.Sprintf("%gkg", k) }

func PToK(p Pound) Kilogramme { return Kilogramme(p * 0.45359237) }
func KToP(k Kilogramme) Pound { return Pound(k / 0.45359237) }
