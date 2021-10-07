package main

import (
	"TestFight/internal/fight"
	"fmt"
	"math/rand"
	"time"
)

func main() {

	const FightDuration = 100
	var winner fight.Fighter

	fighter1 := fight.Boxer{
		fight.Sportsman{"Вася", 200, 120, 10, 113, 15},
	}

	fighter2 := fight.Karateka{
		fight.Sportsman{"Коля", 175, 65, 8, 82, 25},
		25,
	}

	first, second := fight.GetFirstSecond(&fighter1, &fighter2)

	for c := 0; c < FightDuration; c++ {

		rand.Seed(time.Now().UnixNano())
		winner = fight.GetWinner(first, second)

		fmt.Println(first, second)
		if winner != nil {
			break
		}

		first, second = second, first
	}

	fmt.Println("-------------------------------")
	fmt.Println("\u001b[32m", "Winner is ", winner, "\u001b[0m")
}
