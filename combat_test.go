package main

import (
	"testing"
	"log"
)

func TestCombat1(t *testing.T) {
	var m Map
	
	ATTACKRADIUS2 = 4

	m.InitFromString(0,  `
		%%%%%%%%
		%......%
		%.a..b.%
		%....b.%
		%%%%%%%%
	`)
	log.Printf("m: \n%s", &m)

}
