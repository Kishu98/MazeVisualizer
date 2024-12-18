package main

import (
	"fmt"
)

const (
	Cyan = "\033[36m"
)

func VisualizeMaze(maze Maze, visited map[Point]bool, path []Point) {
    fmt.Print("\033[2J\033[H")
	for y, row := range maze.Grid {
		for x, char := range row {
			point := Point{y, x}
			if containsPoint(path, point) {
				fmt.Print(Red, ".", Reset)
			} else if visited[point] {
				fmt.Print(Cyan, ".", Reset)
			} else {
				switch char {
				case 'S':
					fmt.Print(Green, string(char), Reset)
				case 'E':
					fmt.Print(Red, string(char), Reset)
				default:
					fmt.Print(string(char))
				}
			}
		}

		fmt.Println()
	}

}

func containsPoint(path []Point, p Point) bool {
	for _, point := range path {
		if point == p {
			return true
		}
	}

	return false
}
