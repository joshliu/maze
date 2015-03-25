package maze

import (
	"math/rand"
	"fmt"
	)

type Cell struct {
	coordinates []int
	top         []int
	bottom      []int
	right       []int
	left        []int
	north       int
	south       int
	east        int
	west        int
}

type Maze struct {
	count int
	xSize int
	ySize int
	maze  [][]Cell
}

func NewCell(coordinate []int) *Cell {
	c := new(Cell)
	c.coordinates = coordinate
	c.top = []int{coordinate[0], coordinate[1] - 1}
	c.bottom = []int{coordinate[0], coordinate[1] + 1}
	c.right = []int{coordinate[0] + 1, coordinate[1]}
	c.left = []int{coordinate[0] - 1, coordinate[1]}
	c.east = 1
	c.west = 1
	c.north = 1
	c.south = 1
	return c
}

func NewMaze(x int, y int) *Maze {
	m := new(Maze)
	m.xSize = x
	m.ySize = y
	m.count = x * y
	for i := 0; i < y; i++ {
		row := []Cell{}
		for j := 0; j < x; j++ {
			row = append(row, *NewCell([]int{j, i}))
		}
		m.maze = append(m.maze, row)
	}
	return m
}

func Generate(maze Maze, s int64) Maze {
	visited := *new([][]int)
	cellStack := *new([][]int)

	s1 := rand.NewSource(s)
  r := rand.New(s1)

	startingCoords := []int{r.Intn(maze.xSize), r.Intn(maze.ySize)}


	visited = append(visited, startingCoords)
	cellStack = append(cellStack, startingCoords)

	for len(visited) < maze.count {
		//index out of range
		choices := getNeighbors(cellStack[len(cellStack)-1], visited, maze)
		if len(choices) > 0 {
			choice := choices[r.Intn(len(choices))]
			cell1 := maze.maze[cellStack[len(cellStack)-1][1]][cellStack[len(cellStack)-1][0]]
			cell2 := maze.maze[choice[1]][choice[0]]

			if compareSlices(cell1.top, cell2.coordinates) == true {
				cell1.north = 0
				cell2.south = 0
			}
			if compareSlices(cell1.bottom, cell2.coordinates) == true {
				cell1.south = 0
				cell2.north = 0
			}
			if compareSlices(cell1.right, cell2.coordinates) == true {
				cell1.east = 0
				cell2.west = 0
			}
			if compareSlices(cell1.left, cell2.coordinates) == true {
				cell1.west = 0
				cell2.east = 0
			}

			maze.maze[cellStack[len(cellStack)-1][1]][cellStack[len(cellStack)-1][0]] = cell1
			maze.maze[choice[1]][choice[0]] = cell2
			
			visited = append(visited, choice)
			cellStack = append(cellStack, choice)
		} else {
			cellStack = cellStack[:len(cellStack)-1]
		}
	}
	return maze
}

func getNeighbors(coordinates []int, visited [][]int, maze Maze) [][]int {
	x := coordinates[0]
	y := coordinates[1]
	options := *new([][]int)
	left := []int{x - 1, y}
	right := []int{x + 1, y}
	top := []int{x, y - 1}
	bottom := []int{x, y + 1}

	if x-1 >= 0 && !inSlice(left, visited) {
		options = append(options, left)
	}
	if x+1 < maze.xSize && !inSlice(right, visited) {
		options = append(options, right)
	}
	if y-1 >= 0 && !inSlice(top, visited) {
		options = append(options, top)
	}
	if y+1 < maze.ySize && !inSlice(bottom, visited) {
		options = append(options, bottom)
	}
	return options
}

func compareSlices(a []int, b []int) bool {
	if a[0] == b[0] && a[1] == b[1] {
		return true
	}
	return false
}

func inSlice(coords []int, array [][]int) bool {
	for _, b := range array {
		if compareSlices(coords, b) == true {
			return true
		}
	}
	return false
}

func PrintMaze(maze Maze) {
	string := "+"
	for _ = range maze.maze[0] {
		string += "---+"
	}
	string += "\n"
	for _, a := range maze.maze {
		verticalWalls := ""
		bottomWalls := "+"

		for _, b := range a {
			if b.west == 1 {
				verticalWalls += "|"
			}
			if b.west == 0 {
				verticalWalls += " "
			}
			verticalWalls += "   "
			if b.south == 1 {
				bottomWalls += "---+"
			}
			if b.south == 0 {
				bottomWalls += "   +"
			}
		}
		verticalWalls += "|\n"
		bottomWalls += "\n"

		string += verticalWalls
		string += bottomWalls
	}
	string += "\n"

	fmt.Printf(string)
}
