package main

import (
	"os"
	"strconv"
	"strings"
	"fmt"
	"log"
)

//Bot interface defines what we need from a bot
type Bot interface {
	DoTurn(s *State) os.Error
}

//State keeps track of everything we need to know about the state of the game
type State struct {
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
func (s *State) Start() {
	s.Load()
	s.Map = NewMap(s.Rows, s.Cols, s.ViewRadius2)
}

func (s *State) Load() {
	for words := range(getPairs()) {

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

//Loop handles the majority of communication between your bot and the server.
//b's DoWork function gets called each turn after the map has been setup
//BetweenTurnWork gets called after a turn but before the map is reset. It is
//meant to do debugging work.
func (s *State) Loop(b Bot, BetweenTurnWork func()) os.Error {

	//indicate we're ready
	os.Stdout.Write([]byte("go\n"))

	for line := range(getLinesUntil("end")) {
		if line == "go" {
			b.DoTurn(s)

			//end turn
			s.endTurn()

			BetweenTurnWork()

			s.Map.Reset()
			continue
		}

		words := strings.SplitN(line, " ", 5)
		if len(words) < 2 {
			log.Panicf("Invalid command format: \"%s\"", line)
		}

		if words[0] == "turn" {
		turn, _ := strconv.Atoi(words[1])
		if turn != s.Turn+1 {
			log.Panicf("Turn number out of sync, expected %v got %v", s.Turn+1, turn)
		}
		s.Turn = turn
		}
		s.Map.Update(words)
	}

	return nil
}

//Call IssueOrderRowCol to issue an order for an ant at (Row, Col)
func (s *State) IssueOrderRowCol(Row, Col int, d Direction) {
	loc := s.Map.FromRowCol(Row, Col)
	dest := s.Map.Move(loc, d)
	s.Map.RemoveDestination(loc)
	s.Map.AddDestination(dest)
	fmt.Fprintf(os.Stdout, "o %d %d %s\n", Row, Col, d)
}

//Call IssueOrderLoc to issue an order for an ant at loc
func (s *State) IssueOrderLoc(loc Location, d Direction) {
	Row, Col := s.Map.FromLocation(loc)
	dest := s.Map.Move(loc, d)
	s.Map.RemoveDestination(loc)
	s.Map.AddDestination(dest)
	fmt.Fprintf(os.Stdout, "o %d %d %s\n", Row, Col, d)
}

//endTurn is called by Loop, you don't need to call it.
func (s *State) endTurn() {
	os.Stdout.Write([]byte("go\n"))
}
