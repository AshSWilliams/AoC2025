package day4

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func safeCheckIndex(rowIx, columnIx int, paperMap [][]string) int {
	if rowIx >= 0 && rowIx <= len(paperMap)-1 {
		if columnIx >= 0 && columnIx <= len(paperMap[rowIx])-1 {
			if paperMap[rowIx][columnIx] == "@" {
				return 1
			}
		}
	}
	return 0
}

func checkAdjacent(maxAdjacent, rowIx, columnIx int, paperMap [][]string) bool {
	adjacent := safeCheckIndex(rowIx-1, columnIx-1, paperMap) +
		safeCheckIndex(rowIx-1, columnIx, paperMap) +
		safeCheckIndex(rowIx-1, columnIx+1, paperMap) +
		safeCheckIndex(rowIx, columnIx-1, paperMap) +
		safeCheckIndex(rowIx, columnIx+1, paperMap) +
		safeCheckIndex(rowIx+1, columnIx-1, paperMap) +
		safeCheckIndex(rowIx+1, columnIx, paperMap) +
		safeCheckIndex(rowIx+1, columnIx+1, paperMap)
	//fmt.Printf("Paper in position %d, %d has %d adjacent papers\n", rowIx, columnIx, adjacent)
	return adjacent <= maxAdjacent
}

func countPaper(paperMap [][]string) (int, [][]string) {
	newMap := make([][]string, len(paperMap))
	copy(newMap, paperMap)
	total := 0
	maxAdjacent := 3
	for rowIx, row := range paperMap {
		for columnIx, entry := range row {
			if entry == "." {
				continue
			}
			if checkAdjacent(maxAdjacent, rowIx, columnIx, paperMap) {
				//fmt.Printf("Paper in position: %d, %d is accessible\n", rowIx, columnIx)
				total++
				newMap[rowIx][columnIx] = "."
			} else {
				//fmt.Printf("Paper in position: %d, %d is not accessible\n", rowIx, columnIx)
			}
		}
	}
	fmt.Printf("New map: %v", newMap)
	return total, newMap
}

func Main() {
	file, err := os.Open("day4/input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	paperMap := [][]string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		paperMap = append(paperMap, strings.Split(row, ""))
	}

	fmt.Printf("Read map: %v\n", paperMap)

	accessible, _ := countPaper(paperMap)

	fmt.Printf("Total accaccessible: %d\n", accessible)

}

func Main2() {
	file, err := os.Open("day4/input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	paperMap := [][]string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		paperMap = append(paperMap, strings.Split(row, ""))
	}

	total := 0

	done := false
	for !done {
		removed, newMap := countPaper(paperMap)
		if removed == 0 {
			done = true
		}
		total += removed
		paperMap = newMap
	}
	fmt.Printf("Total removed: %d", total)
}
