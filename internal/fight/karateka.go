package fight

import (
	"fmt"
)

type Karateka struct {
	Sportsman
	DamageLeg int
}

//func (k Karateka) Hit() int {
//	//if rand.Intn(2) > 0 {
//	return k.MainProperties.DamageFist
//	//}
//	//return k.MainProperties.DamageLeg
//}

func (k *Karateka) GetDamage(d int) {
	k.Health -= d
}

func (k Karateka) String() string {
	return fmt.Sprintf("%v (%v health)", k.Name, k.Health)
}
