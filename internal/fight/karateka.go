package fight

import (
	"fmt"
	"math/rand"
)

type Karateka struct {
	Name string
	Height int
	Weight float32
	DamageFist int
	DamageLeg int
	Health int
	EvasionProbability int
}

func (k *Karateka) Hit() int {
	if rand.Intn(2) > 0 {
		return k.DamageFist
	}
	return k.DamageLeg
}

func (k *Karateka) Dodge() bool {
	return rand.Intn(100) < k.EvasionProbability
}

func (k *Karateka) GetDamage(d int) {
	k.Health -= d
}

func (k *Karateka) CanFight() bool {
	return k.Health > 0
}

func (k Karateka) String() string {
	return fmt.Sprintf("%v (%v height, %v weight)", k.Name, k.Height, k.Weight)
}
