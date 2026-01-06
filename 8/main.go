package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/stephen-condon/advent-of-code-2025/utilities"
)

type Point struct {
	x, y, z int
}

type Edge struct {
	i, j     int
	distance float64
}

type UnionFind struct {
	parent []int
	size   []int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	return &UnionFind{parent: parent, size: size}
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x]) // Path compression
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) bool {
	rootX := uf.Find(x)
	rootY := uf.Find(y)

	if rootX == rootY {
		return false // Already in same set
	}

	// Union by size
	if uf.size[rootX] < uf.size[rootY] {
		uf.parent[rootX] = rootY
		uf.size[rootY] += uf.size[rootX]
	} else {
		uf.parent[rootY] = rootX
		uf.size[rootX] += uf.size[rootY]
	}
	return true
}

func (uf *UnionFind) GetComponentSizes() []int {
	sizeMap := make(map[int]int)
	for i := 0; i < len(uf.parent); i++ {
		root := uf.Find(i)
		sizeMap[root] = uf.size[root]
	}

	sizes := make([]int, 0, len(sizeMap))
	for _, size := range sizeMap {
		sizes = append(sizes, size)
	}
	return sizes
}

func main() {
	fmt.Println("Part One:")
	executePartOne("example.txt", 10)
	executePartOne("input.txt", 1000)

	fmt.Println("\nPart Two:")
	executePartTwo("example.txt")
	executePartTwo("input.txt")
}

func executePartOne(filename string, connections int) {
	input := utilities.LoadInput(filename)
	if len(input) == 0 {
		fmt.Println("No input found")
		return
	}

	result := solveJunctionBoxes(input, connections)
	fmt.Printf("%s: %d\n", filename, result)
}

func executePartTwo(filename string) {
	input := utilities.LoadInput(filename)
	if len(input) == 0 {
		fmt.Println("No input found")
		return
	}

	result := findUnifyingConnection(input)
	fmt.Printf("%s: %d\n", filename, result)
}

func findUnifyingConnection(lines []string) int {
	// Parse points
	points := make([]Point, 0, len(lines))
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		parts := strings.Split(line, ",")
		if len(parts) != 3 {
			continue
		}
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])
		points = append(points, Point{x, y, z})
	}

	n := len(points)

	// Calculate all pairwise distances
	edges := make([]Edge, 0, n*(n-1)/2)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dist := distance(points[i], points[j])
			edges = append(edges, Edge{i, j, dist})
		}
	}

	// Sort edges by distance
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].distance < edges[j].distance
	})

	// Use Union-Find and connect edges until all are in one circuit
	uf := NewUnionFind(n)

	for _, edge := range edges {
		if uf.Union(edge.i, edge.j) {
			// Check if all nodes are now in one circuit
			// This happens when we have exactly 1 component
			sizes := uf.GetComponentSizes()
			if len(sizes) == 1 {
				// This connection unified everything
				return points[edge.i].x * points[edge.j].x
			}
		}
	}

	return 0
}

func solveJunctionBoxes(lines []string, connections int) int {
	// Parse points
	points := make([]Point, 0, len(lines))
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		parts := strings.Split(line, ",")
		if len(parts) != 3 {
			continue
		}
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])
		points = append(points, Point{x, y, z})
	}

	n := len(points)

	// Calculate all pairwise distances
	edges := make([]Edge, 0, n*(n-1)/2)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dist := distance(points[i], points[j])
			edges = append(edges, Edge{i, j, dist})
		}
	}

	// Sort edges by distance
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].distance < edges[j].distance
	})

	// Use Union-Find to connect closest pairs
	uf := NewUnionFind(n)

	for i := 0; i < connections && i < len(edges); i++ {
		uf.Union(edges[i].i, edges[i].j)
	}

	// Get component sizes
	sizes := uf.GetComponentSizes()
	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))

	fmt.Printf("  Circuit sizes: %v\n", sizes)

	// Multiply three largest
	if len(sizes) < 3 {
		return 0
	}
	return sizes[0] * sizes[1] * sizes[2]
}

func distance(p1, p2 Point) float64 {
	dx := float64(p1.x - p2.x)
	dy := float64(p1.y - p2.y)
	dz := float64(p1.z - p2.z)
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}
