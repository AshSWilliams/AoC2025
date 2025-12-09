package day9

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

type rectangle struct {
	p1 point
	p2 point
}

func between(x1, x2, p int) bool {
	if x1 < x2 {
		return p > x1 && p < x2
	} else {
		return p < x1 && p > x2
	}
}

func (p point) pointInRectangle(r rectangle) bool {
	return between(r.p1.x, r.p2.x, p.x) && between(r.p1.y, r.p2.y, p.y)
}

func (p point) rectangleSize(q point) int {
	xmax := max(p.x, q.x)
	xmin := min(p.x, q.x)
	ymax := max(p.y, q.y)
	ymin := min(p.y, q.y)
	return (xmax - xmin + 1) * (ymax - ymin + 1)
}

func (r rectangle) rectangleSize() int {
	return r.p1.rectangleSize(r.p2)
}

func (p point) isEqual(q point) bool {
	return p.x == q.x && p.y == q.y
}

type line struct {
	p1 point
	p2 point
}

func (l line) pointsInLine() []point {
	points := []point{}
	// vertical
	if l.p1.x == l.p2.x {
		upper := max(l.p1.y, l.p2.y)
		lower := min(l.p1.y, l.p2.y)
		for y := lower; y <= upper; y++ {
			points = append(points, point{x: l.p1.x, y: y})
		}
	} else {
		// Horizontal
		upper := max(l.p1.x, l.p2.x)
		lower := min(l.p1.x, l.p2.x)
		for x := lower; x <= upper; x++ {
			points = append(points, point{x: x, y: l.p1.y})
		}
	}
	return points
}

// Whether a line goes *inside* a rectangle
func (l line) lineWithinRectangle(r rectangle) bool {
	for _, p := range l.pointsInLine() {
		if p.pointInRectangle(r) {
			return true
		}
	}
	return false
}

type polygon struct {
	points map[int]point
	sides  map[int]line
}

func (pg *polygon) addPoint(p point, ix int) {
	pg.points[ix] = p
	if len(pg.points) > 1 {
		pg.sides[ix] = line{p1: pg.points[ix-1], p2: p}
	}

}

func (pg *polygon) close() {
	sideNo := len(pg.sides)
	pg.sides[sideNo] = line{p1: pg.points[sideNo], p2: pg.points[0]}
}

func (pg polygon) rectangleIntersects(r rectangle) bool {
	for _, side := range pg.sides {
		if side.lineWithinRectangle(r) {
			return true
		}
	}
	return false
}

func Main() {
	file, err := os.Open("day9/input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	coordMap := map[int]point{}

	scanner := bufio.NewScanner(file)
	lineNo := 0
	for scanner.Scan() {
		coords := strings.Split(scanner.Text(), ",")
		x, err := strconv.Atoi(coords[0])
		if err != nil {
			log.Fatalf("Failed to parse coordinate: %v", err)
		}
		y, err := strconv.Atoi(coords[1])
		if err != nil {
			log.Fatalf("Failed to parse coordinate: %v", err)
		}
		coordMap[lineNo] = point{x: x, y: y}
		lineNo++
	}
	fmt.Printf("Found coords: %v\n", coordMap)
	maxSize := 0
	for _, point1 := range coordMap {
		for _, point2 := range coordMap {
			if !point2.isEqual(point1) {
				size := point1.rectangleSize(point2)
				if size > maxSize {
					maxSize = size
				}
			}
		}
	}
	fmt.Printf("Largest found rectangle: %d\n", maxSize)
}

func Main2() {
	file, err := os.Open("day9/input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	pg := polygon{
		points: map[int]point{},
		sides:  map[int]line{},
	}

	scanner := bufio.NewScanner(file)
	lineNo := 0
	for scanner.Scan() {
		coords := strings.Split(scanner.Text(), ",")
		x, err := strconv.Atoi(coords[0])
		if err != nil {
			log.Fatalf("Failed to parse coordinate: %v", err)
		}
		y, err := strconv.Atoi(coords[1])
		if err != nil {
			log.Fatalf("Failed to parse coordinate: %v", err)
		}
		pg.addPoint(point{x: x, y: y}, lineNo)
		lineNo++
	}
	pg.close()

	//fmt.Printf("Parsed polygon: %v\n", pg)

	// Find rectangles, check if they're inside the polygon
	rectangles := []rectangle{}
	for ix, point1 := range pg.points {
		for jx, point2 := range pg.points {
			if ix > jx {
				rectangles = append(rectangles, rectangle{
					p1: point1,
					p2: point2,
				})
			}
		}
	}
	// Sort in descending order
	sort.Slice(rectangles, func(i, j int) bool {
		return rectangles[i].rectangleSize() > rectangles[j].rectangleSize()
	})

	//fmt.Printf("Sorted rectangles: %v\n", rectangles)

	biggestRectangle := rectangle{}

	for _, rectangle := range rectangles {
		if !pg.rectangleIntersects(rectangle) {
			//fmt.Printf("Rectangle does not intersect: %v\n", rectangle)
			biggestRectangle = rectangle
			break
		} else {
			//fmt.Printf("Rectangle does intersect: %v\n", rectangle)
		}
	}

	fmt.Printf("Found biggest rectangle: %v with area: %d\n", biggestRectangle, biggestRectangle.rectangleSize())

}
