package main

import (
	"aoc/day1"
	"flag"
)

func main() {
	puzzleDay := flag.Int("day", 1, "Day to solve")
	switch *puzzleDay {
	case 1:
		day1.Main()
	}
}
