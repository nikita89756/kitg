package main

import (
	"fmt"
)

type Point struct {
	Name          string
	PointChildren []*Point
}

type Graph struct {
	points []*Point
}

func (g *Graph) AddPoint(name string) *Point {
	point := &Point{Name: name}
	g.points = append(g.points, point)
	return point
}

func (g *Graph) AddEdge(parent, child *Point) {
	parent.PointChildren = append(parent.PointChildren, child)
	child.PointChildren = append(child.PointChildren, parent)
}

func (g *Graph) DFS(start *Point, visited map[*Point]bool, component []*Point) []*Point {
	stack := []*Point{start}
	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if !visited[current] {
			visited[current] = true
			component = append(component, current)
			for _, child := range current.PointChildren {
				if !visited[child] {
					stack = append(stack, child)
				}
			}
		}
	}
	return component
}

func (g *Graph) FindConnectedComponents() [][]*Point {
	visited := make(map[*Point]bool)
	var components [][]*Point

	for _, point := range g.points {
		if !visited[point] {
			var component []*Point
			component = g.DFS(point, visited, component)
			components = append(components, component)
		}
	}

	return components
}

func main() {

	graph := &Graph{}

	graph1 := &Graph{}
	pointA1 := graph1.AddPoint("A1")
	pointB1 := graph1.AddPoint("B1")
	pointC1 := graph1.AddPoint("C1")
	graph1.AddEdge(pointA1, pointB1)
	graph1.AddEdge(pointB1, pointC1)
	graph1.AddEdge(pointC1, pointA1)

	graph2 := &Graph{}
	pointA2 := graph2.AddPoint("A2")
	pointB2 := graph2.AddPoint("B2")
	pointC2 := graph2.AddPoint("C2")
	pointD2 := graph2.AddPoint("D2")
	graph2.AddEdge(pointA2, pointB2)
	graph2.AddEdge(pointC2, pointD2)

	graph3 := &Graph{}
	pointA3 := graph3.AddPoint("A3")
	pointB3 := graph3.AddPoint("B3")
	_ = graph3.AddPoint("C3")
	graph3.AddEdge(pointA3, pointB3)

	graph5 := &Graph{}
	pointA5 := graph5.AddPoint("A5")
	pointB5 := graph5.AddPoint("B5")
	pointC5 := graph5.AddPoint("C5")
	graph5.AddEdge(pointA5, pointB5)
	graph5.AddEdge(pointB5, pointC5)
	graph5.AddEdge(pointC5, pointA5)

	pointA := graph.AddPoint("A")
	pointB := graph.AddPoint("B")
	pointC := graph.AddPoint("C")
	pointD := graph.AddPoint("D")
	pointE := graph.AddPoint("E")
	_ = graph.AddPoint("F")

	graph.AddEdge(pointA, pointB)
	graph.AddEdge(pointB, pointC)
	graph.AddEdge(pointD, pointE)

	components := graph.FindConnectedComponents()

	for i, component := range components {
		fmt.Printf("Component %d:\n", i+1)
		for _, point := range component {
			fmt.Println(point.Name)
		}
		fmt.Println()
	}
}
