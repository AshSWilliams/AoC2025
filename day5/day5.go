package day5

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Map[T, V any](ts []T, fn func(T) V) []V {
	result := make([]V, len(ts))
	for i, t := range ts {
		result[i] = fn(t)
	}
	return result
}

type singleRange struct {
	lower int64
	upper int64
}

func (s singleRange) withinRange(i int64) bool {
	return i >= s.lower && i <= s.upper
}

func rangeFromText(s string) *singleRange {
	bounds := strings.Split(s, "-")
	lower, err := strconv.ParseInt(bounds[0], 10, 64)
	if err != nil {
		log.Fatalf("Failed to create range from input: %v", err)
	}
	upper, err := strconv.ParseInt(bounds[1], 10, 64)
	if err != nil {
		log.Fatalf("Failed to create range from input: %v", err)
	}
	return &singleRange{
		lower: int64(lower),
		upper: int64(upper),
	}
}

func (s singleRange) toInterval() []int64 {
	return []int64{s.lower, s.upper}
}

type multiRange struct {
	ranges []*singleRange
}

func (m *multiRange) addRange(s *singleRange) {
	m.ranges = append(m.ranges, s)
}

func (m multiRange) withinAnyRange(i int64) bool {
	for _, idRange := range m.ranges {
		if idRange.withinRange(i) {
			//fmt.Printf("%d is within range %v\n", i, idRange)
			return true
		}
	}
	return false
}

func (m *multiRange) mergeRanges() [][]int64 {
	originalRanges := Map(m.ranges, func(r *singleRange) []int64 { return r.toInterval() })
	sort.Slice(originalRanges, func(i, j int) bool {
		return int(originalRanges[i][0]) < int(originalRanges[j][0])
	})

	merged := [][]int64{}
	currentInterval := originalRanges[0]
	merged = append(merged, currentInterval)

	for _, interval := range originalRanges {
		currentEnd := currentInterval[1]
		nextStart := interval[0]
		nextEnd := interval[1]

		if currentEnd >= nextStart {
			currentInterval[1] = max(currentEnd, nextEnd)
		} else {
			currentInterval = interval
			merged = append(merged, currentInterval)
		}
	}

	return merged
}

func Main() {
	file, err := os.Open("day5/input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	ranges := multiRange{}
	totalFresh := 0

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.Contains(line, "-") {
			newRange := rangeFromText(line)
			ranges.addRange(newRange)
		} else if line != "" {
			id, err := strconv.ParseInt(line, 10, 64)
			if err != nil {
				log.Fatalf("Failed to read ID: %v", err)
			}
			if ranges.withinAnyRange(int64(id)) {
				totalFresh++
			}
		}
	}
	fmt.Printf("Total of %d spoiled items\n", totalFresh)

	// Part 2
	var rangeSize int64
	rangeSize = 0
	mergedRanges := ranges.mergeRanges()
	for _, intervals := range mergedRanges {
		rangeSize += intervals[1] - intervals[0] + 1
	}

	fmt.Printf("Total range sizes: %d\n", rangeSize)
}
