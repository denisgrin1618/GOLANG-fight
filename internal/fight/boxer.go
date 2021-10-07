package fight

import (
	"fmt"
)

type Boxer struct {
	Sportsman
}

func (b *Boxer) GetDamage(d int) {
	b.Health -= d
}
func (b Boxer) String() string {
	return fmt.Sprintf("%v (%v health)", b.Name, b.Health)
}
