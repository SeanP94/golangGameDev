package main

import (
	"fmt"
	"time"
)

type GameState int

const (
	loading GameState = iota
	menu
	pause
	world
	dungeon
	exit
)

var statename = map[GameState]string{
	loading: "loading",
	menu:    "menu",
	pause:   "pause",
	world:   "world",
	dungeon: "dungeon",
	exit:    "exit",
}

func (gs *GameState) currentState() {

}

func testrun() {
	lastUpdate := time.Now()
	currState := loading
	fmt.Println("Current Game State :: " + statename[currState])
	for currState != exit {
		currTime := time.Now()
		if currTime.Sub(lastUpdate).Seconds() >= 2 {
			currState++
			fmt.Println("Current Game State :: " + statename[currState])
			lastUpdate = currTime
		}
		time.Sleep(1 * time.Second)
	}
}
