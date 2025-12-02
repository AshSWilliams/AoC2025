package day2

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type idRange struct {
	firstID int
	lastID  int
	invalid []int
}

func newIDRange(strRange string) (*idRange, error) {
	first, err := strconv.Atoi(strings.Split(strRange, "-")[0])
	if err != nil {
		return nil, err
	}
	last, err := strconv.Atoi(strings.Split(strRange, "-")[1])
	if err != nil {
		return nil, err
	}
	return &idRange{
		firstID: first,
		lastID:  last,
		invalid: make([]int, 0),
	}, nil
}

func isInvalid(x string) bool {
	length := len(x)
	if length%2 != 0 {
		return false
	} else {
		return x[:length/2] == x[length/2:]
	}
}

func findFactors(x int) []int {
	factors := make([]int, 0)
	for i := 1; i < x; i++ {
		if x%i == 0 {
			factors = append(factors, i)
		}
	}
	return factors
}

func isInvalid2(x string) bool {
	length := len(x)
	factors := findFactors(length)
	for i := 0; i < len(factors); i++ {
		factor := factors[i]
		// n equal pieces
		pieces := make([]string, 0)
		for j := range length / factor {
			pieces = append(pieces, x[j*factor:(j+1)*factor])
		}
		if allSameStrings(pieces) {
			return true
		}
	}
	return false
}

func allSameStrings(x []string) bool {
	for i := 1; i < len(x); i++ {
		if x[i] != x[0] {
			return false
		}
	}
	return true
}

func (i *idRange) findInvalids(part2 bool) {
	// easiest way to iterate between integers
	for x := i.firstID; x <= i.lastID; x++ {
		strx := strconv.Itoa(x)
		if part2 {
			if isInvalid2(strx) {
				i.invalid = append(i.invalid, x)
			}
		} else {
			if isInvalid(strx) {
				i.invalid = append(i.invalid, x)
			}
		}
	}
}

func Main(part2 bool) {
	file := "day2/input"
	input, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("Failed to open input: %v", err)
	}

	total := 0

	strInput := string(input[:])
	ranges := strings.Split(strInput, ",")
	for _, strRange := range ranges {
		idRange, err := newIDRange(strRange)
		if err != nil {
			log.Fatalf("Failed to parse ID range: %v", err)
		}
		idRange.findInvalids(part2)
		for _, invalidID := range idRange.invalid {
			total += invalidID
		}
	}
	fmt.Printf("Total of invalid ids: %d \n", total)
}
