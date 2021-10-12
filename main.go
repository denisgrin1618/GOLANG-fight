package main

import (
	"TestFight/internal/fight"
	"fmt"
)

func main() {

	message1 := make(chan string)
	message2 := make(chan string)
	message3 := make(chan string)
	message4 := make(chan string)

	fighter1 := fight.Boxer{
		fight.Sportsman{"Вася", 200, 120, 10, 100, 15},
	}

	fighter2 := fight.Karateka{
		fight.Sportsman{"Коля", 175, 65, 8, 0, 25},
		25,
	}

	fighter3 := fight.Boxer{
		fight.Sportsman{"Петя", 200, 120, 10, 50, 15},
	}

	fighter4 := fight.Boxer{
		fight.Sportsman{"Женя", 200, 120, 10, 80, 15},
	}

	ring := fight.NewRing(100)
	ring.AddFighter(&fighter1)
	ring.AddFighter(&fighter2)
	ring.AddFighter(&fighter3)
	ring.AddFighter(&fighter4)

	go ring.StartFight(&fighter1, message1)
	go ring.StartFight(&fighter2, message2)
	go ring.StartFight(&fighter3, message3)
	go ring.StartFight(&fighter4, message4)

	for {
		select {
		case <-ring.Done:
			fmt.Println("-------------------------------")
			fmt.Println("\u001b[32m", "Winner is ", ring.GetWinner(), "\u001b[0m")
			return
		case msg1 := <-message1:
			fmt.Println(msg1)
		case msg2 := <-message2:
			fmt.Println(msg2)
		case msg3 := <-message3:
			fmt.Println(msg3)
		case msg4 := <-message4:
			fmt.Println(msg4)
		}
	}

	ring.Wait()
}
