package weightconv

import "fmt"

type Pound float64
type Kilogramme float64

func (p Pound) String() string      { return fmt.Sprintf("%g°C", p) }
func (k Kilogramme) String() string { return fmt.Sprintf("%g°F", k) }
