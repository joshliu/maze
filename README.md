MAZE
====
Generate a Maze!

Install using:
```
  $ go get github.com/joshliu/maze
```
Then run go install in the maze directory.
```
  $ go install
```

It takes three optional arguments for width, length, and seed.
```
  $ maze [width] [length] [seed]
```

Mazes of the same size and same seed will be the same. <br />
The default seed value is 0, and default size is 10x10