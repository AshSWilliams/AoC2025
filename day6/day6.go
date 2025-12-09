package day6

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type column struct {
	numbers   []int
	operation string
}

func (c column) solve() int {
	if c.operation == "+" {
		sum := 0
		for _, number := range c.numbers {
			sum += number
		}
		return sum
	} else if c.operation == "*" {
		prod := 1
		for _, number := range c.numbers {
			prod *= number
		}
		return prod
	}
	log.Fatal("Unexpected!")
	return 0
}

type multiColumn struct {
	columns []*column
}

func (m *multiColumn) addNumberLine(tokens []string) {
	for ix, token := range tokens {
		intToken, err := strconv.Atoi(token)
		if err != nil {
			log.Fatalf("Failed to read line: %v", err)
		}
		if ix > len(m.columns)-1 {
			m.columns = append(m.columns, &column{numbers: []int{intToken}})
		} else {
			m.columns[ix].numbers = append(m.columns[ix].numbers, intToken)
		}
	}
}

func (m *multiColumn) addNumberLines2(tokens [][]string, problems int) {
	// Find columns which divide problems
	dividers := []int{-1}
	for ix := range tokens[0] {
		dividingLine := true
		for j := range tokens {
			if tokens[j][ix] != " " {
				dividingLine = false
			}
		}
		if dividingLine {
			dividers = append(dividers, ix)
		}
	}
	dividers = append(dividers, len(tokens[0]))
	for ix := range problems {
		startChar := dividers[ix] + 1
		endChar := dividers[ix+1]
		fmt.Printf("Considering from columns %d to %d\n", startChar, endChar)
		for i := startChar; i < endChar; i++ {
			strNumber := tokens[0][i] + tokens[1][i] + tokens[2][i] + tokens[3][i]
			number, err := strconv.Atoi(strings.TrimSpace(strNumber))
			if err != nil {
				log.Fatalf("Failed to parse: %v", err)
			}
			if ix > len(m.columns)-1 {
				m.columns = append(m.columns, &column{numbers: []int{number}})
			} else {
				m.columns[ix].numbers = append(m.columns[ix].numbers, number)
			}
		}
	}
}

func (m *multiColumn) addOperationLine(tokens []string) {
	for ix, token := range tokens {
		if token != "*" && token != "+" {
			log.Fatalf("Unexpected operation: %s", token)
		}
		m.columns[ix].operation = token
	}
}

func (m *multiColumn) solve() int {
	total := 0
	for _, problem := range m.columns {
		total += problem.solve()
	}
	return total
}

func Main() {
	file, err := os.Open("day6/input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	columns := multiColumn{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Fields(scanner.Text())
		if !strings.Contains(line, "+") {
			columns.addNumberLine(tokens)
		} else {
			columns.addOperationLine(tokens)
		}
	}
	fmt.Printf("Total: %d\n", columns.solve())
}

func Main2() {
	file, err := os.Open("day6/input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	columns := multiColumn{}
	lines := [][]string{}
	opLine := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.Contains(line, "+") {
			lines = append(lines, strings.Split(line, ""))
		} else {
			opLine = strings.Fields(line)
		}
	}
	problems := len(opLine)
	columns.addNumberLines2(lines, problems)
	columns.addOperationLine(opLine)
	fmt.Printf("Total: %d\n", columns.solve())
}
