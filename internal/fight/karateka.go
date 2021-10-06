package fight

import (
	"fmt"
	"math/rand"
)

type Karateka struct {
	MainProperties Sportsman
	DamageLeg      int
}

func (k Karateka) Hit() int {
	//if rand.Intn(2) > 0 {
	return k.MainProperties.DamageFist
	//}
	//return k.MainProperties.DamageLeg
}

func (k Karateka) Dodge() bool {
	return rand.Intn(100) < k.MainProperties.EvasionProbability
}

func (k *Karateka) GetDamage(d int) {
	k.MainProperties.Health -= d
}

func (k Karateka) CanFight() bool {
	return k.MainProperties.Health > 0
}

func (k Karateka) String() string {
	return fmt.Sprintf("%v (%v health)", k.MainProperties.Name, k.MainProperties.Health)
}
