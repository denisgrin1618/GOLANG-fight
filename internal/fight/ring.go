package fight

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Ring struct {
	AllFighters   []Fighter
	mutex         sync.Mutex
	Done          chan bool
	FightDuration int
}

func NewRing(FightDuration int) Ring {
	return Ring{[]Fighter{}, sync.Mutex{}, make(chan bool), FightDuration}
}

func (ring *Ring) AddFighter(f Fighter) {
	ring.AllFighters = append(ring.AllFighters, f)
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

func (ring *Ring) StartFight(fighter Fighter, message chan string, wg *sync.WaitGroup) {

	for i := 0; i < ring.FightDuration; i++ {
		if ring.GetWinner() != nil {
			ring.Done <- true
			return
		}

		if !fighter.CanFight() {
			wg.Done()
			message <- fmt.Sprintf("%s, lost", fighter)
			return
		}

		ring.RemoveFightersWhoCanNotFight()
		opponent := ring.GetRandomOpponentFighter(fighter)
		if opponent != nil && !opponent.Dodge() {
			opponent.GetDamage(fighter.Hit())
			message <- fmt.Sprintf("%s -> %s", fighter, opponent)
		}

	}

}
