package main

import (
	"fmt"
	"os"
	"strings"
)

type node struct {
	name      string
	edges     map[*edge]*node
	travelled bool
}

func newNode(name string) *node { return &node{name: name, edges: make(map[*edge]*node)} }

type edge struct {
	travelled bool
}

type graph struct {
	nodes map[string]*node
	edges []*edge
}

func (g *graph) resetNodes() {
	for _, n := range g.nodes {
		n.travelled = false
	}
}

func (g *graph) resetEdges() {
	for _, e := range g.edges {
		e.travelled = false
	}
}

func (g *graph) removeShortestPath(source, dest *node) bool {
	type queueItem struct {
		e    *edge
		n    *node
		prev *queueItem
	}

	queue := make([]*queueItem, 0, len(g.nodes))
	queue = append(queue, &queueItem{n: source})

	found := false
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr.n == dest {
			for itr := curr; itr.e != nil; itr = itr.prev {
				itr.e.travelled = true
			}
			found = true
			break
		}

		for e, n := range curr.n.edges {
			if e.travelled || n.travelled {
				continue
			}
			n.travelled = true
			queue = append(queue, &queueItem{e, n, curr})
		}
	}
	g.resetNodes()
	return found
}

func (g *graph) cutPaths(source, dest *node, pathNum int) bool {
	complete := true
	for i := 0; i < pathNum; i++ {
		if !g.removeShortestPath(source, dest) {
			complete = false
			break
		}
	}
	g.resetEdges()
	return complete
}

func (g *graph) split(cuts int) ([]*node, []*node) {
	g1, g2 := []*node{}, []*node{}

	var source *node
	for _, n := range g.nodes {
		source = n
		break
	}
	g1 = append(g1, source)

	for _, dest := range g.nodes {
		if source == dest {
			continue
		}

		if g.cutPaths(source, dest, cuts+1) {
			g1 = append(g1, dest)
		} else {
			g2 = append(g2, dest)
		}
	}
	return g1, g2
}

func parseInput(inputData []byte) *graph {

	input := strings.Split(string(inputData), "\n")

	nodes := make(map[string]*node)
	for _, line := range input {
		name, _, _ := strings.Cut(line, ": ")
		nodes[name] = newNode(name)
	}
	edges := make([]*edge, 0)
	for _, line := range input {
		sourceName, destNames, _ := strings.Cut(line, ": ")
		source := nodes[sourceName]
		for _, destName := range strings.Split(destNames, " ") {
			dest, ok := nodes[destName]
			if !ok {
				dest = newNode(destName)
				nodes[destName] = dest
			}
			newEdge := &edge{}
			edges = append(edges, newEdge)
			source.edges[newEdge] = dest
			dest.edges[newEdge] = source
		}
	}
	return &graph{nodes, edges}
}

func Part1(inputData []byte) int {

	g := parseInput(inputData)
	g1, g2 := g.split(3)

	return len(g1) * len(g2)
}

func main() {

	pwd, _ := os.Getwd()
	inputData, err := os.ReadFile(pwd + "/2023/day25/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(Part1(inputData))
	//fmt.Println(Part2(inputData))
}
