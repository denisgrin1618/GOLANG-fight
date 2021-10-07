package fight

import "math/rand"

type Sportsman struct {
	Name               string
	Height             int
	Weight             float32
	DamageFist         int
	Health             int
	EvasionProbability int
}

func (s Sportsman) Hit() int {
	return s.DamageFist
}

func (s Sportsman) Dodge() bool {
	return rand.Intn(100) < s.EvasionProbability
}

func (s Sportsman) CanFight() bool {
	return s.Health > 0
}
