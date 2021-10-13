package fight

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Ring struct {
	AllFighters []Fighter
	mutex       sync.Mutex
	stop        chan bool
	CountRounds int
}

func NewRing(CountRounds int) Ring {
	return Ring{
		[]Fighter{},
		sync.Mutex{},
		make(chan bool),
		CountRounds,
	}
}

func (ring *Ring) AddFighter(f Fighter) {
	ring.AllFighters = append(ring.AllFighters, f)
}

func (ring *Ring) RemoveFighter(fighterRemove Fighter) {

	ring.mutex.Lock()

	index, ok := ring.GetIndexFighter(fighterRemove)
	if ok {
		ring.AllFighters[index] = ring.AllFighters[len(ring.AllFighters)-1]
		ring.AllFighters = ring.AllFighters[:len(ring.AllFighters)-1]
	}

	ring.mutex.Unlock()
}

func (ring *Ring) GetIndexFighter(fighterRemove Fighter) (int, bool) {

	for i, fighter := range ring.AllFighters {
		if fighterRemove == fighter {
			return i, true
		}
	}

	return -1, false
}

func (ring *Ring) RemoveFightersWhoCanNotFight() {

	ring.mutex.Lock()

	newF := make([]Fighter, 0)
	for i, fighter := range ring.AllFighters {
		if fighter.CanFight() {
			newF = append(newF, ring.AllFighters[i])
		}
	}
	ring.AllFighters = newF

	ring.mutex.Unlock()
}

func (ring *Ring) GetRandomOpponentFighter(exceptionFighter Fighter) Fighter {

	if len(ring.AllFighters) == 0 {
		return nil
	}

	if len(ring.AllFighters) == 1 && ring.AllFighters[0] == exceptionFighter {
		return nil
	}

	rand.Seed(time.Now().UnixNano())

	for {

		ring.mutex.Lock()
		randomFighter := ring.AllFighters[rand.Intn(len(ring.AllFighters))]
		ring.mutex.Unlock()

		if randomFighter != exceptionFighter {
			return randomFighter
		}

	}

}

func (ring *Ring) GetWinner() Fighter {

	if len(ring.AllFighters) == 1 {
		return ring.AllFighters[0]
	}
	return nil
}

func (ring *Ring) StartFight(fighter Fighter) {

	for {
		select {
		case <-ring.stop:
			return

		case currentRound := <-fighter.GetCurrentRound():

			if !fighter.CanFight() {
				ring.RemoveFighter(fighter)
				fmt.Printf("Round: %d   %s lost \n", currentRound, fighter)
				return
			}

			opponent := ring.GetRandomOpponentFighter(fighter)
			if opponent != nil {
				opponent.GetDamage(fighter.Hit())
				fmt.Printf("Round: %d   %s -> %s \n", currentRound, fighter, opponent)
			}

		}
	}
}

func (ring *Ring) StartCountRounds() {

	defer ring.stopFight()

	for currentRound := 1; currentRound <= ring.CountRounds; currentRound++ {

		//added timeout for sequential output
		<-time.After(100 * time.Millisecond)

		for i := 0; i < len(ring.AllFighters); i++ {
			ring.AllFighters[i].GetCurrentRound() <- currentRound
		}

		if ring.GetWinner() != nil {
			return
		}
	}

}

func (ring *Ring) stopFight() {
	ring.stop <- true
	fmt.Println("-------------------------------")
	fmt.Println("\u001b[32m", "Winner is ", ring.GetWinner(), "\u001b[0m")
}
