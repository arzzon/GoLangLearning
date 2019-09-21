package main

import "fmt"

// vertex represents a single vertex in a graph
type vertex struct {
	ID        string        // ID is the name of the vertex
	Neighbors *neighborList // Neighbors is the list of connected vertices/nodes
}

// newVertex creates a new default Vertex
func newVertex(id string) *vertex {
	return &vertex{
		ID:        id,
		Neighbors: newNeighborList(),
	}
}

// neighbor reoresents a single vertice/node
type neighbor struct {
	ID   string
	Next *neighbor // Stores the reference of the next neighbor
}

// newNeighbor creates a new default Neighbor
func newNeighbor(id string) *neighbor {
	return &neighbor{
		ID:   id,
		Next: nil,
	}
}

// neighborList represents a list of neighbors
type neighborList struct {
	Start *neighbor // Start points to the Start of the NeighborList
	Size  int       // Stores the total number of neighbors
}

// newNeighborList creates a new default NeighborList
func newNeighborList() *neighborList {
	return &neighborList{
		Start: nil,
		Size:  0,
	}
}

// insertNeighbor adds a new neighbor to the neighbor list
func (nl *neighborList) insertNeighbor(id string) {
	// Check if no neighbor exist
	if nl.Start == nil {
		nl.Start = newNeighbor(id)
		nl.Size++
	} else {
		tempNode := nl.Start
		for tempNode.Next != nil && tempNode.ID != id {
			tempNode = tempNode.Next
		}
		if tempNode.ID != id {
			tempNode.Next = newNeighbor(id)
			nl.Size++
		} else {
			fmt.Println("Neighbor already exists!")
		}
	}
}

// showNeighbors dispalys all the neighbors in a neighbor list
func (nl *neighborList) showNeighbors() {
	tempNode := nl.Start
	for tempNode != nil {
		fmt.Printf("%s ", tempNode.ID)
		tempNode = tempNode.Next
	}
	fmt.Println()
}

// removeNeighbor removes the node/neighbor from the neighbor list
func (nl *neighborList) removeNeighbor(id string) {
	tempNode1 := nl.Start
	tempNode2 := tempNode1
	for tempNode2 != nil && tempNode2.ID != id {
		tempNode1 = tempNode2
		tempNode2 = tempNode1.Next
	}
	// No node/neighbor exists
	if tempNode1 == nil {
		fmt.Println("NeighborList is empty!")
	} else if tempNode2 == nil {
		// Node node found
	} else {
		// Remove the node
		tempNode1.Next = tempNode2.Next
		tempNode2.Next = nil
		nl.Size--
	}
}

// NewGraph creates a new default graph
func NewGraph(dir ...bool) *Graph {
	if len(dir) != 0 && dir[0] == true {
		return &Graph{
			Vertices:  make(map[string]*vertex),
			NVertices: 0,
			Directed:  true,
		}
	}
	return &Graph{
		Vertices:  make(map[string]*vertex),
		NVertices: 0,
		Directed:  false,
	}
}

// Graph will have a set of nodes/vertices and edges
type Graph struct {
	Vertices  map[string]*vertex
	NVertices int
	Directed  bool
}

// AddVertex adds a new vertex to the graph
func (g *Graph) AddVertex(id string) {
	// Check whether node already exists
	if _, ok := g.Vertices[id]; ok {
		fmt.Printf("Vertex: %s already exists.", id)
	} else {
		g.Vertices[id] = newVertex(id)
		g.NVertices++
	}
}

// AddEdge adds an edge between to nodes/vertices
func (g *Graph) AddEdge(src string, dst string) {
	//Check whether src & dst vertices exist or not
	vSrc, ok1 := g.Vertices[src]
	vDst, ok2 := g.Vertices[dst]
	if ok1 && ok2 {
		vSrc.Neighbors.insertNeighbor(dst)
		if !g.Directed { // If undirected graph then add the edge from dst -> src as well
			vDst.Neighbors.insertNeighbor(src)
		}
	} else {
		fmt.Printf("Failed to add edge as some of the vertex(s) do(es)n't exist.")
	}
}

// ShowGraph shows the entire graph with the edges
func (g *Graph) ShowGraph() {
	for vertexK, vertexV := range g.Vertices {
		fmt.Printf("%s: ", vertexK)
		vertexV.Neighbors.showNeighbors()
	}
}

func main() {
	//g := NewGraph(true)
	//g := NewGraph(false)
	g := NewGraph()
	g.AddVertex("A")
	g.AddVertex("B")
	g.AddVertex("C")
	g.AddVertex("D")
	g.AddVertex("E")
	g.AddVertex("F")
	g.AddEdge("A", "B")
	g.AddEdge("B", "C")
	g.AddEdge("C", "D")
	g.AddEdge("D", "E")
	g.AddEdge("E", "A")
	g.AddEdge("B", "E")
	g.ShowGraph()
}
