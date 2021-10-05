package main

import (
	"TestFight/internal/fight"
	"fmt"
)

func main() {

	const FightDuration = 100
	var winner fight.Fighter

	fighter1 := fight.Boxer{
		"Вася",
		200,
		120,
		10,
		113,
		15,
	}

	fighter2 := fight.Karateka{
		"Коля",
		175,
		65,
		8,
		25,
		82,
		25,
	}

	for c := 0; c < FightDuration; c++ {

		fight.Round(&fighter1, &fighter2)
		if !fighter1.CanFight() {
			winner = &fighter2
			break
		}

		fight.Round(&fighter2, &fighter1)
		if !fighter2.CanFight() {
			winner = &fighter1
			break
		}
	}

	fmt.Println("Winner is ", winner)
}
