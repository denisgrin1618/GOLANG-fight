package fight

import (
	"math/rand"
	"time"
)

type Fighter interface {
	Hit() int
	Dodge() bool
	GetDamage(int)
	CanFight() bool
}

func GetWinner(f1 Fighter, f2 Fighter) Fighter {
	if !f2.Dodge() {
		f2.GetDamage(f1.Hit())
	}

	if !f2.CanFight() {
		return f1
	}

	return nil
}

func GetFirstSecond(f1 Fighter, f2 Fighter) (Fighter, Fighter) {

	rand.Seed(time.Now().Unix())
	first := f1
	second := f2
	for i := 0; i < rand.Intn(10); i++ {
		first, second = second, first
	}

	return first, second
}
