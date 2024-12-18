package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Maze struct {
	Grid  [][]rune
	Start [2]int
	End   [2]int
}

const (
    Reset = "\033[0m"
    Green = "\033[32m"
    Red = "\033[31m"
    Yellow = "\033[33m"
    )

func main() {
	file, err := os.OpenFile("Input.txt", os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input string
	for scanner.Scan() {
		line := scanner.Text()
		input += line + "\n"
	}

	maze, err := parseMaze(input)

    fmt.Println("Initial Maze:")
    DisplayMaze(maze)

    path, err := Dijkstra(maze)
    if err != nil {
        log.Fatal("Error solving maze:", err)
        return
    }

    fmt.Println("Shortest path:", path)
}

func parseMaze(input string) (Maze, error) {
    var grid [][]rune
    var start [2]int
    var end [2]int
    startFound, endFound := false, false

    lines := strings.Split(input, "\n")
    for r, line := range lines {
        if strings.TrimSpace(line) == "" {
            continue
        }
        row := []rune(line) 
        for c, cell := range row {
            if cell == 'S' {
                if startFound {
                    return Maze{}, fmt.Errorf("Multiple start positions found.")
                }
                start = [2]int{r, c}
                startFound = true
            }
            if cell == 'E' {
                if endFound {
                    return Maze{}, fmt.Errorf("Multiple end positions found")
                }
                end = [2]int{r, c}
                endFound = true
            }
        }
        grid = append(grid, row)
    }

    if !startFound {
        return Maze{}, fmt.Errorf("No start position found")
    }
    if !endFound {
        return Maze{}, fmt.Errorf("No end position found")
    }

    return Maze{Grid: grid, Start: start, End: end},nil
}

func DisplayMaze(maze Maze) {
    for _, row := range maze.Grid {
        for _, cell := range row {
            switch cell {
            case 'S':
            fmt.Print(Green, string(cell), Reset)
            case 'E':
            fmt.Print(Red, string(cell), Reset)
            default:
            fmt.Print(string(cell))
            }
        }
        fmt.Println()
    }
}
