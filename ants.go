package main

import (
	"strconv"
	"strings"
	"fmt"
	"log"
)

//Game keeps track of everything we need to know about the state of the game
type Game struct {
	LoadTime      int   //in milliseconds
	TurnTime      int   //in milliseconds
	Rows          int   //number of rows in the map
	Cols          int   //number of columns in the map
	Turns         int   //maximum number of turns in the game
	ViewRadius2   int   //view radius squared
	AttackRadius2 int   //battle radius squared
	SpawnRadius2  int   //spawn radius squared
	PlayerSeed    int64 //random player seed
	Turn          int   //current turn number

	Map *Map
}

//Start takes the initial parameters from stdin
func (s *Game) Start() {
	s.Load()
	s.Map = NewMap(s.Rows, s.Cols, s.ViewRadius2)
}

func (s *Game) Load() {
	for words := range getPairs() {

		param, _ := strconv.Atoi(words[1])

		switch words[0] {
		case "loadtime":
			s.LoadTime = param
		case "turntime":
			s.TurnTime = param
		case "rows":
			s.Rows = param
		case "cols":
			s.Cols = param
		case "turns":
			s.Turns = param
		case "viewradius2":
			s.ViewRadius2 = param
		case "attackradius2":
			s.AttackRadius2 = param
		case "spawnradius2":
			s.SpawnRadius2 = param
		case "player_seed":
			param64, _ := strconv.Atoi64(words[1])
			s.PlayerSeed = param64
		case "turn":
			s.Turn = param

		default:
			log.Printf("unknown command: %v", words)
		}
	}
}

//main initializes the state and starts the processing loop
func main() {
	var (
		g Game
		bot MyBot
	)
	g.Start()

	//indicate we're ready
	fmt.Println("go")

	for line := range getLinesUntil("end") {
		if line == "go" {
			bot.DoTurn(&g)

			//end turn
			fmt.Println("go")

			g.Map.Reset()
			continue
		}

		words := strings.SplitN(line, " ", 5)
		if len(words) < 2 {
			log.Panicf("Invalid command format: \"%s\"", line)
		}

		if words[0] == "turn" {
			turn, _ := strconv.Atoi(words[1])
			if turn != g.Turn+1 {
				log.Panicf("Turn number out of sync, expected %v got %v", g.Turn+1, turn)
			}
			g.Turn = turn
		} else {
			g.Map.Update(words)
		}
	}
}
