package main

import (
	"bufio"
	"os"
	"log"
	"strings"
	"strconv"
)

var stdin = bufio.NewReader(os.Stdin)

// Read pairs of words from stdin until we read single word 'ready'.
func getPairs() chan []string {
	ch := make(chan []string)
	go func() {
		for line := range getLinesUntil("ready") {
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

// Read lines from stdin until we get end string
func getLinesUntil(end string) chan string {
	ch := make(chan string)
	go func() {
		for {
			line, err := stdin.ReadString('\n')
			if err != nil {
				log.Panicln(err)
			}
			line = line[:len(line)-1] //remove the delimiter

			if line == "" {
				continue
			}

			if line == end {
				break
			}

			ch <- line
		}
		close(ch)
	}()
	return ch
}

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
func atoi64(s string) int64 {
	i, _ := strconv.Atoi64(s)
	return i
}
