package day1

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type safe struct {
	dialPos   int
	zeros     int
	crossings int
}

func newSafe() safe {
	safe := safe{
		dialPos:   50,
		zeros:     0,
		crossings: 0,
	}
	return safe
}

func mod(a, b int) int {
	return (a%b + b) % b
}

func (s *safe) rotate(direction string, clicks int) {
	old := s.dialPos

	step := 1
	if direction == "L" {
		step = -1
	}

	// Count crossings
	pos := old
	for range clicks {
		pos = mod(pos+step, 100)
		if pos == 0 {
			s.crossings++
		}
	}

	// Final position
	s.dialPos = pos
	if s.dialPos == 0 {
		s.zeros++
	}

	log.Printf("Rotated %s clicks %d, new position: %d, zeros %d",
		direction, clicks, s.dialPos, s.zeros)
}

func Main() {
	file, err := os.Open("day1/input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	safe := newSafe()

	// Read input
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineSlice := strings.Split(line, "")
		clicks, err := strconv.Atoi(strings.Join(lineSlice[1:], ""))
		if err != nil {
			log.Fatalf("Failed to read clicks from line %s: %v", line, err)
		}
		safe.rotate(lineSlice[0], clicks)
	}
	log.Printf("Safe: %v", safe)
}
