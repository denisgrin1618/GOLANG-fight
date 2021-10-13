package main

import (
	"TestFight/internal/fight"
)

func main() {

	fighter1 := fight.Boxer{
		fight.Sportsman{"Вася", 200, 120, 10, 100, 15, make(chan int)},
	}
	fighter2 := fight.Karateka{
		fight.Sportsman{"Коля", 175, 65, 8, 120, 25, make(chan int)},
		25,
	}
	fighter3 := fight.Boxer{
		fight.Sportsman{"Петя", 200, 120, 10, 100, 15, make(chan int)},
	}
	fighter4 := fight.Boxer{
		fight.Sportsman{"Женя", 200, 120, 10, 90, 15, make(chan int)},
	}

	ring := fight.NewRing(100)
	ring.AddFighter(&fighter1)
	ring.AddFighter(&fighter2)
	ring.AddFighter(&fighter3)
	ring.AddFighter(&fighter4)

	go ring.StartFight(&fighter1)
	go ring.StartFight(&fighter2)
	go ring.StartFight(&fighter3)
	go ring.StartFight(&fighter4)

	ring.StartCountRounds()

}
