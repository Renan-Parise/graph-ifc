package usecases

import (
	"container/heap"
	"errors"

	"github.com/Renan-Parise/graph-ifc/entities"
)

type PathFinder struct {
	Graph *entities.Graph
}

func NewPathFinder(graph *entities.Graph) *PathFinder {
	return &PathFinder{Graph: graph}
}

func (pf *PathFinder) FindShortestPath(from, to string) ([]string, float64, error) {
	graph := pf.Graph

	if _, ok := graph.Nodes[from]; !ok {
		return nil, 0, errors.New("source node not found")
	}
	if _, ok := graph.Nodes[to]; !ok {
		return nil, 0, errors.New("destination node not found")
	}

	distances := make(map[string]float64)
	previous := make(map[string]string)
	for node := range graph.Nodes {
		distances[node] = -1
	}
	distances[from] = 0

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{Value: from, Priority: 0})

	for pq.Len() > 0 {
		currentItem := heap.Pop(pq).(*Item)
		current := currentItem.Value
		if current == to {
			break
		}

		for _, edge := range graph.Edges[current] {
			alt := distances[current] + edge.Distance
			if distances[edge.To] == -1 || alt < distances[edge.To] {
				distances[edge.To] = alt
				previous[edge.To] = current
				heap.Push(pq, &Item{Value: edge.To, Priority: alt})
			}
		}
	}

	path := []string{}
	u := to
	if _, ok := previous[u]; !ok && u != from {
		return nil, 0, errors.New("no path found")
	}
	for u != "" {
		path = append([]string{u}, path...)
		u = previous[u]
	}

	return path, distances[to], nil
}

func (pf *PathFinder) FindAllPaths(from, to string) ([][]string, []float64, error) {
	graph := pf.Graph

	if _, ok := graph.Nodes[from]; !ok {
		return nil, nil, errors.New("source node not found")
	}
	if _, ok := graph.Nodes[to]; !ok {
		return nil, nil, errors.New("destination node not found")
	}

	var paths [][]string
	var costs []float64

	var dfs func(current string, visited map[string]bool, path []string, cost float64)
	dfs = func(current string, visited map[string]bool, path []string, cost float64) {
		visited[current] = true
		path = append(path, current)

		if current == to {
			pathCopy := make([]string, len(path))
			copy(pathCopy, path)
			paths = append(paths, pathCopy)
			costs = append(costs, cost)
		} else {
			for _, edge := range graph.Edges[current] {
				if !visited[edge.To] {
					dfs(edge.To, visited, path, cost+edge.Distance)
				}
			}
		}

		visited[current] = false
		path = path[:len(path)-1]
	}

	dfs(from, make(map[string]bool), []string{}, 0.0)

	if len(paths) == 0 {
		return nil, nil, errors.New("no paths found")
	}

	return paths, costs, nil
}

type Item struct {
	Value    string
	Priority float64
	Index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.Index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *Item, value string, priority float64) {
	item.Value = value
	item.Priority = priority
	heap.Fix(pq, item.Index)
}
