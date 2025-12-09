package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func highestFromSubstring(ss []string) (int, int) {
	highest := 1
	highestIx := 0
	for ix, digit := range ss {
		intDigit, err := strconv.Atoi(digit)
		if err != nil {
			log.Fatalf("Failed to parse battery: %v", err)
		}
		if intDigit > highest && ix < len(ss) {
			highest = intDigit
			highestIx = ix
		}
	}
	return highest, highestIx + 1
}

// Main algorithm
func findHighestFromBattery(battery []string, N int) int {
	// Call len(battery) = M, joltageLen = N
	switchedOn := ""
	ix := 0
	M := len(battery)
	// Search for highest digit in range ix to (M - (N - i)), where i is the iteration number
	for i := range N {
		substringToCheck := battery[ix : M-(N-i)+1]
		highest, ixOffset := highestFromSubstring(substringToCheck)
		// Found highest at an index, next iteration search from that index
		ix = ix + ixOffset
		switchedOn += strconv.Itoa(highest)
	}
	joltage, err := strconv.Atoi(switchedOn)
	if err != nil {
		log.Fatalf("Failed to convert joltage %s", switchedOn)
	}
	return joltage
}

func Main(joltageLen int) {
	file, err := os.Open("day3/input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	total := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		battery := strings.Split(scanner.Text(), "")
		joltage := findHighestFromBattery(battery, joltageLen)
		total += joltage
	}
	fmt.Printf("Total: %d\n", total)
}
