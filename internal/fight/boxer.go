package fight

import (
	"fmt"
	"math/rand"
)

type Boxer struct {
	MainProperties Sportsman
}

func (b Boxer) Hit() int {
	return b.MainProperties.DamageFist
}

func (b Boxer) Dodge() bool {
	return rand.Intn(100) < b.MainProperties.EvasionProbability
}

func (b *Boxer) GetDamage(d int) {
	b.MainProperties.Health -= d
}

func (b Boxer) CanFight() bool {
	return b.MainProperties.Health > 0
}

func (b Boxer) String() string {
	return fmt.Sprintf("%v (%v health)", b.MainProperties.Name, b.MainProperties.Health)
}
