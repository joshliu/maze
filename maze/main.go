package main

import (
	"github.com/joshliu/maze"
	"os"
	"strconv"
)

func main() {
	x, y := 0,0
	if len(os.Args) < 2 {
		x, y = 10, 10
	} else if len(os.Args) < 3 {
		x, _ = strconv.Atoi(os.Args[1])
		y, _ =  strconv.Atoi(os.Args[1])
	} else {
		x, _ = strconv.Atoi(os.Args[1])
		y, _ = strconv.Atoi(os.Args[2])
	}
	lol := maze.Generate(*maze.NewMaze(x,y))
	maze.PrintMaze(lol)
}