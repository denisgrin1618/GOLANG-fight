package fight

import (
	"fmt"
	"math/rand"
)

type Boxer struct {
	Name               string
	Height             int
	Weight             float32
	DamageFist         int
	Health             int
	EvasionProbability int
}

func (b *Boxer) Hit() int {
	return b.DamageFist
}

func (b *Boxer) Dodge() bool {
	return rand.Intn(100) < b.EvasionProbability
}

func (b *Boxer) GetDamage(d int) {
	b.Health -= d
}

func (b *Boxer) CanFight() bool {
	return b.Health > 0
}

func (b Boxer) String() string {
	//return fmt.Sprintf("%v (%v height, %v weight)", b.Name, b.Height, b.Weight)
	return fmt.Sprintf("%v (%v height, %v weight)", b.Name, b.Height, b.Weight)
}
