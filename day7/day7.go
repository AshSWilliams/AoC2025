package day7

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func splitter(beams map[int]int, splitterLoc int) (map[int]int, bool) {
	val, ok := beams[splitterLoc]
	if ok {
		beams[splitterLoc-1] += val
		beams[splitterLoc+1] += val
		delete(beams, splitterLoc)
	}
	return beams, ok
}

func countSplits(manifold [][]string) (int, int) {
	splits := 0
	beams := map[int]int{}
	for _, manifoldLine := range manifold {
		for ix, char := range manifoldLine {
			switch char {
			case ".":
				continue
			case "S":
				beams[ix] = 1
			case "^":
				var split bool
				beams, split = splitter(beams, ix)
				if split {
					splits += 1
				}
			}
		}
	}
	paths := 0
	for _, v := range beams {
		paths += v
	}
	return splits, paths
}

func Main() {
	file, err := os.Open("day7/input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lines := [][]string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		lines = append(lines, line)
	}
	for _, line := range lines {
		fmt.Printf("Parsed line %v\n", line)
	}

	splits, paths := countSplits(lines)

	fmt.Printf("Found %d splits\n", splits)
	fmt.Printf("Found %d paths\n", paths)
}
