package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Point struct {
	y, x int
}

type PriorityQueueItem struct {
	Point    Point
	Distance int
	Index    int
}

type PriorityQueue []PriorityQueueItem

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Distance < pq[j].Distance
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(PriorityQueueItem)
	item.Index = len(*pq)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.Index = -1
	*pq = old[:n-1]
	return item
}

func Dijkstra(maze Maze) ([]Point, error) {
	rows, cols := len(maze.Grid), len(maze.Grid[0])
	dist := make([][]int, rows)
	prev := make(map[Point]Point)
    visited := make(map[Point]bool)
	directions := []Point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	for i := range dist {
		dist[i] = make([]int, cols)
		for j := range dist[i] {
			dist[i][j] = math.MaxInt32
		}
	}

	start := Point{maze.Start[0], maze.Start[1]}
	end := Point{maze.End[0], maze.End[1]}
	dist[start.y][start.x] = 0

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, PriorityQueueItem{Point: start, Distance: 0})

	for pq.Len() > 0 {
		current := heap.Pop(pq).(PriorityQueueItem)
		currPoint := current.Point

        visited[currPoint] = true
        // VisualizeMaze(maze, visited, nil)

		if currPoint == end {
			break
		}

		for _, dir := range directions {
			neighbour := Point{currPoint.y + dir.y, currPoint.x + dir.x}
			if neighbour.y < 0 || neighbour.x < 0 || neighbour.y >= rows || neighbour.x >= cols {
				continue
			}
			if maze.Grid[neighbour.y][neighbour.x] == '#' {
				continue
			}

			newDist := dist[currPoint.y][currPoint.x] + 1
			if newDist < dist[neighbour.y][neighbour.x] {
				dist[neighbour.y][neighbour.x] = newDist
				prev[neighbour] = currPoint
				heap.Push(pq, PriorityQueueItem{Point: neighbour, Distance: newDist})
			}
		}
	}

    path := []Point{}
    for at := end; at != start; at = prev[at] {
        path = append([]Point{at}, path...)
        if _, found := prev[at]; !found {
            return nil, fmt.Errorf("No path found")
        }
    }
    path = append([]Point{start}, path...)

    for i := 0; i <= len(path); i++ {
        VisualizeMaze(maze, visited, path[:i])
    }

    return path, nil
}
