package main

import (
	"strings"
	"fmt"
	"log"
	"os"
)

type Turn uint

//Game keeps track of everything we need to know about the state of the game
type Game struct {
	LoadTime     int   //in milliseconds
	TurnTime     int   //in milliseconds
	Rows         int   //number of rows in the map
	Cols         int   //number of columns in the map
	Turns        int   //maximum number of turns in the game
	ViewRadius2  int   //view radius squared
	SpawnRadius2 int   //spawn radius squared
	PlayerSeed   int64 //random player seed
}

func (s *Game) Load() {
	for words := range getPairs() {

		param := atoi(words[1])

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
		case "turn":
			TURN = Turn(param)
		case "viewradius2":
			s.ViewRadius2 = param
		case "attackradius2":
			ATTACKRADIUS2 = param
			log.Printf("ATTACKRADIUS2: %d", param)
		case "spawnradius2":
			s.SpawnRadius2 = param
		case "player_seed":
			param64 := atoi64(words[1])
			s.PlayerSeed = param64
		default:
			log.Printf("unknown command loading game: %v", words)
		}
	}
	log.Printf("Game stats: %+v", *s)
}

func nsec() int64 {
	s, ns, _ := os.Time()
	return s*1e9 + ns
}

type Timer struct {
	start int64
	split int64
}

func (t *Timer) Reset() {
	now := nsec()
	t.start = now
	t.split = now
}

func (t *Timer) Split(s string) {
	now := nsec()

	deltaSplit := float64(now-t.split) / 1e9
	delta := float64(now-t.start) / 1e9
	t.split = now

	if true {
		log.Printf("%s: %.3f %.3f", s, deltaSplit, delta)
	}
}

//main initializes the state and starts the processing loop
func main() {
	var (
		g Game
		m Map
	)
	g.Load()
	m.Init(g.Rows, g.Cols, g.ViewRadius2)

	// Think time is fraction of turn time
	// (and converting milliseconds to nanoseconds)
	m.thinkTime = int64(float64(g.TurnTime) * 0.9e6)

	//indicate we're ready
	fmt.Println("go")

	var t Timer
	for line := range getLinesUntil("end") {
		if line == "go" {
			m.setDeadline()
			t.Reset()
			m.UpdatesProcessed()
			t.Split("updates")
			m.DoTurn(&t)
			t.Split("turn")

			//end turn
			fmt.Println("go")

			m.Reset()
			t.Split("reset")
			continue
		}

		words := strings.SplitN(line, " ", 5)
		if len(words) < 2 {
			log.Panicf("Invalid command format: \"%s\"", line)
		}

		m.Update(words)
	}
}
