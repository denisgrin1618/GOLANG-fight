package main

import (
	"TestFight/internal/fight"
	"fmt"
)

func main() {

	var winner fight.Fighter
	
	fighter1 := fight.Boxer{"Boxer", 184, 89.9, 100, 1000, 50}
	fighter2 := fight.Karateka{"Karateka", 184, 89.9, 100, 120, 1000, 30}

	FightDuration := 100
	for c:=0; c < FightDuration; c++ {

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
