package main

import (
	"fmt"
	"math/rand"
)

type door struct {
	ID      int
	Winning bool
	Chosen  bool
	Open    bool
}

const (
	doorCount      = 3
	iterationCount = 1000000
)

func main() {
	var doors [doorCount]door
	var stayedWin int
	var stayedLose int
	var switchedWin int
	var switchedLose int

	for range iterationCount {
		// Initialize all doors
		for i := range doors {
			doors[i].ID = i + 1
			doors[i].Winning = false
			doors[i].Chosen = false
			doors[i].Open = false
		}

		// Pick a door that will have the prize
		doors[rand.Intn(doorCount)].Winning = true

		// ** Player picks a door
		playerFirstChoice := &doors[rand.Intn(len(doors))]
		playerFirstChoice.Chosen = true // Mark it as being chosen in the first guess

		// ** Host opens one of the remaining doors
		// Find out the remaining doors (not chosen)
		remainingDoors := make([]*door, 0, doorCount-1)
		for i := range doors {
			if doors[i].Chosen {
				continue
			}
			remainingDoors = append(remainingDoors, &doors[i])
		}

		// Of those doors filter out the winning one
		doorsThatCanBeOpened := make([]*door, 0)
		for i := range remainingDoors {
			if remainingDoors[i].Winning {
				continue
			}
			doorsThatCanBeOpened = append(doorsThatCanBeOpened, remainingDoors[i])
		}
		// Pick one of the doors to open
		hostOpenedDoor := doorsThatCanBeOpened[rand.Intn(len(doorsThatCanBeOpened))]
		hostOpenedDoor.Open = true

		// ** Player chooses a door with or without switching
		var playerSecondChoice *door
		for i := range doors {
			if doors[i].Open || doors[i].Chosen {
				continue
			}
			playerSecondChoice = &doors[i]
		}

		if playerSecondChoice == nil {
			panic("playerSecondChoice is nil")
		}

		playerSwitched := rand.Intn(2) == 1

		if !playerSwitched {
			if playerFirstChoice.Winning {
				// Player WON by NOT switching the door!
				stayedWin++
			} else {
				// Player LOST by NOT switching the door!
				stayedLose++
			}
		} else {
			if playerSecondChoice.Winning {
				// Player WON by switching the door!
				switchedWin++
			} else {
				// Player LOST by switching the door!
				switchedLose++
			}
		}
	}

	fmt.Printf("stayedWin: %v\n", stayedWin)
	fmt.Printf("stayedLose: %v\n", stayedLose)
	fmt.Printf("switchedWin: %v\n", switchedWin)
	fmt.Printf("switchedLose: %v\n", switchedLose)
}
