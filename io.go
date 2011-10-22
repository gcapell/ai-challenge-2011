package main

import (
	"bufio"
	"os"
	"log"
	"strings"
)

var stdin = bufio.NewReader(os.Stdin)

// Read pairs of words from stdin until we read single word 'ready'.
func getPairs() (chan []string) {
	ch := make (chan []string)
	go func() {
		for {
			line, err := stdin.ReadString('\n')
			if err != nil {
				log.Panicln( err)
			}
			line = line[:len(line)-1] //remove the delimiter
	
			if line == "" {
				continue
			}
	
			if line == "ready" {
				break
			}
	
			words := strings.SplitN(line, " ", 2)
			if len(words) != 2 {
				log.Panicf("bad pair line: %s", line)
			}
			ch <- words
		}
		close(ch)
	}()

	return ch
}

