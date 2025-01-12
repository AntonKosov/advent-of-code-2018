package main

import (
	"fmt"
	smath "math"
	"strings"

	"github.com/AntonKosov/advent-of-code-2018/aoc/input"
	"github.com/AntonKosov/advent-of-code-2018/aoc/math"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() []rune {
	regex := []rune(input.Lines()[0])
	return regex[1 : len(regex)-1]
}

func process(regex []rune) int {
	maze := NewMaze()
	initPos := math.NewVector2(0, 0)
	route := compoundRoute(&regex)
	route(initPos, nil, maze)

	fmt.Println(maze)

	return farRooms(maze)
}

func farRooms(maze Maze) int {
	passedDoors := 0
	visited := make(map[math.Vector2[int]]bool, len(maze)/2)
	initPos := math.NewVector2(0, 0)
	currentPositions := []math.Vector2[int]{initPos}
	visited[initPos] = true
	dirs := []math.Vector2[int]{{X: -1}, {X: 1}, {Y: -1}, {Y: 1}}
	rooms := 0
	for len(currentPositions) > 0 {
		if passedDoors >= 1000 {
			rooms += len(currentPositions)
		}
		nextPositions := make([]math.Vector2[int], 0, len(currentPositions))
		for len(currentPositions) > 0 {
			pos := currentPositions[0]
			currentPositions = currentPositions[1:]
			for _, dir := range dirs {
				nextPos := pos.Add(dir)
				if maze[nextPos] == DoorCell {
					nextPos = nextPos.Add(dir)
					if !visited[nextPos] {
						visited[nextPos] = true
						nextPositions = append(nextPositions, nextPos)
					}
				}
			}
		}

		currentPositions = nextPositions
		passedDoors++
	}

	return rooms
}

type Cell uint8

const (
	UnawailableCell Cell = iota
	RoomCell        Cell = iota
	DoorCell        Cell = iota
)

const (
	North = 'N'
	East  = 'E'
	South = 'S'
	West  = 'W'
)

var dirRegex = map[rune]bool{North: true, East: true, South: true, West: true}

var dirs = map[rune]math.Vector2[int]{North: {Y: -1}, East: {X: 1}, South: {Y: 1}, West: {X: -1}}

type Maze map[math.Vector2[int]]Cell

func NewMaze() Maze {
	return Maze{math.NewVector2(0, 0): RoomCell}
}

func (m Maze) mergeTo(target Maze, origing math.Vector2[int]) {
	for pos, cell := range m {
		target[pos.Add(origing)] = cell
	}
}

func (m Maze) String() string {
	minX, minY, maxX, maxY := smath.MaxInt, smath.MaxInt, 0, 0
	for pos := range m {
		minX = min(minX, pos.X)
		minY = min(minY, pos.Y)
		maxX = max(maxX, pos.X)
		maxY = max(maxY, pos.Y)
	}

	cellToRune := map[Cell]rune{UnawailableCell: '#', RoomCell: '.', DoorCell: '/'}

	var sb strings.Builder
	for r := minY - 1; r <= maxY+1; r++ {
		for c := minX - 1; c <= maxX+1; c++ {
			sb.WriteRune(cellToRune[m[math.NewVector2(c, r)]])
		}
		sb.WriteRune('\n')
	}

	return sb.String()
}

type Route func(pos math.Vector2[int], followingRoute Route, maze Maze) (endPos math.Vector2[int])

func singleRoute(regex *[]rune) Route {
	var route []rune
	for len(*regex) > 0 && dirRegex[(*regex)[0]] {
		route = append(route, (*regex)[0])
		*regex = (*regex)[1:]
	}

	return func(pos math.Vector2[int], _ Route, maze Maze) math.Vector2[int] {
		for _, dirRune := range route {
			dir := dirs[dirRune]
			pos = pos.Add(dir)
			maze[pos] = DoorCell
			pos = pos.Add(dir)
			maze[pos] = RoomCell
		}

		return pos
	}
}

func compoundRoute(regex *[]rune) Route {
	var routes []Route
	for len(*regex) > 0 {
		start := (*regex)[0]
		if start == '|' || start == ')' {
			break
		}

		routeBuilder := groupRoutes
		if dirRegex[start] {
			routeBuilder = singleRoute
		}

		routes = append(routes, routeBuilder(regex))
	}

	return compoundRouteBuilder(routes)
}

func compoundRouteBuilder(routes []Route) Route {
	var cachedMap Maze
	var cachedPos math.Vector2[int]

	return func(pos math.Vector2[int], _ Route, maze Maze) math.Vector2[int] {
		if cachedMap == nil {
			cachedMap = NewMaze()
			for i, route := range routes {
				cachedPos = route(cachedPos, compoundRouteBuilder(routes[i+1:]), cachedMap)
			}
		}

		cachedMap.mergeTo(maze, pos)

		return pos.Add(cachedPos)
	}
}

func groupRoutes(regex *[]rune) Route {
	optional := false
	var routes []Route
	*regex = (*regex)[1:]
	for (*regex)[0] != ')' {
		routes = append(routes, compoundRoute(regex))
		if (*regex)[0] == '|' {
			*regex = (*regex)[1:]
			optional = (*regex)[0] == ')'
		}
	}
	*regex = (*regex)[1:]

	if !optional {
		return func(pos math.Vector2[int], followingRoute Route, maze Maze) math.Vector2[int] {
			for _, route := range routes {
				route(pos, followingRoute, maze)
			}

			return pos
		}
	}

	return func(pos math.Vector2[int], followingRoute Route, maze Maze) math.Vector2[int] {
		for _, route := range routes {
			nextPos := route(pos, followingRoute, maze)
			followingRoute(nextPos, followingRoute, maze)
		}

		return pos
	}
}
