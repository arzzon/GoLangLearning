package Graph

type Vertex struct {
	ID        string
	Neighbors *NeighborList
}

type Neighbor struct {
	ID   string
	Next *Neighbor
}

type NeighborList struct {
	Start *Neighbor
	Size  int
}

func NewNeighborList() *NeighborList {
	return &NeighborList{
		Start: nil,
		Size:  0,
	}
}

func NewGraph() *Graph {
	return &Graph{
		Vertices:  make(map[string]*Vertex),
		NVertices: 0,
	}
}

func (nl *NeighborList) InsertNeighbor(id string) {

}

func (nl *NeighborList) ShowNeighbors() {

}

type Graph struct {
	Vertices  map[string]*Vertex
	NVertices int
}

func (g *Graph) AddVertex(id string) {}

func (g *Graph) AddEdge(src string, dst string) {}
