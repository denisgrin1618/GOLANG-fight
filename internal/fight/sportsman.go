package fight

import (
	"fmt"
	"math/rand"
)

type Sportsman struct {
	Name               string
	Height             int
	Weight             float32
	DamageFist         int
	Health             int
	EvasionProbability int
	CurrentRound       chan int
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

func (s Sportsman) String() string {
	return fmt.Sprintf("%v (%v health)", s.Name, s.Health)
}

func (s Sportsman) GetCurrentRound() chan int {
	return s.CurrentRound
}
