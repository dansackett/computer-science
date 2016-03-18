package graph

// Vertex represents a single node in a graph.
type Vertex struct {
	// Value in our implementation is just an integer for readability
	Value int
	// Graph allows us to make sure the vertices are in the same graph
	Graph *Graph
	// Neighbors stores edges between two vertices
	Neighbors []*Vertex
}

// GetNeighbors is a convenient function to return the neighboring vertices
func (v *Vertex) GetNeighbors() []*Vertex {
	return v.Neighbors
}

// IsAdjacent checks if the current vertex has an edge with another
func (v *Vertex) IsAdjacent(other *Vertex) bool {
	if v.Graph != other.Graph {
		return false
	}

	for _, n := range v.Neighbors {
		if n == other {
			return true
		}
	}

	return false
}

// RemoveEdge removes a neighboring vertex from its list
func (v *Vertex) RemoveEdge(other *Vertex) bool {
	if v.Graph != other.Graph {
		return false
	}

	if !v.IsAdjacent(other) {
		return false
	}

	updateNeighbors := func(cur, target *Vertex) []*Vertex {
		var removeIndex int
		for i, n := range cur.Neighbors {
			if n == target {
				removeIndex = i
			}
		}

		if len(cur.Neighbors) == 1 {
			return []*Vertex{}
		}
		return append(cur.Neighbors[:removeIndex], cur.Neighbors[removeIndex+1:]...)
	}

	v.Neighbors = updateNeighbors(v, other)
	other.Neighbors = updateNeighbors(other, v)
	return true
}

// Graph is a container which holds a list of vertices within it
type Graph struct {
	Vertices []*Vertex
}

// NewGraph is a convenience method to instantiate a new graph object
func NewGraph() *Graph {
	return &Graph{}
}

// Adjacent checks if two vertices are adjacent to each other
func (g *Graph) Adjacent(first, second *Vertex) bool {
	if first.Graph != g || second.Graph != g {
		return false
	}
	return first.IsAdjacent(second)
}

// Edges returns an edge list showing pairs of vertices that connect
func (g *Graph) Edges() [][]int {
	var edges [][]int

	for _, v := range g.Vertices {
		for _, n := range v.Neighbors {
			edges = append(edges, []int{v.Value, n.Value})
		}
	}

	return edges
}

// AddVertex inserts a new vertex into the graph returning that new vertex
func (g *Graph) AddVertex(value int) *Vertex {
	v := &Vertex{Value: value, Graph: g}
	g.Vertices = append(g.Vertices, v)
	return v
}

// RemoveVertex removes a given vertex from a graph and its neighbor relations
func (g *Graph) RemoveVertex(v *Vertex) bool {
	var removeIndex int

	if v.Graph != g {
		return false
	}

	for i, vertex := range g.Vertices {
		if vertex == v {
			removeIndex = i
			for _, neighbor := range vertex.Neighbors {
				neighbor.RemoveEdge(v)
			}
		}
	}

	g.Vertices = append(g.Vertices[:removeIndex], g.Vertices[removeIndex+1:]...)

	return true
}

// AddEdge creates an edge between two vertices
func (g *Graph) AddEdge(first, second *Vertex) {
	if first.Graph == second.Graph && !first.IsAdjacent(second) {
		second.Neighbors = append(second.Neighbors, first)
		first.Neighbors = append(first.Neighbors, second)
	}
}

// RemoveEdge removes an edge between two vertices if it exists
func (g *Graph) RemoveEdge(first, second *Vertex) {
	if first.Graph == second.Graph && first.IsAdjacent(second) {
		first.RemoveEdge(second)
	}
}

// GetVertex searches for a vertex based on the value
func (g *Graph) GetVertex(value int) *Vertex {
	for _, v := range g.Vertices {
		if v.Value == value {
			return v
		}
	}
	return nil
}

// ToAdjacencyList returns a graph represented as an Adjacency List
func (g *Graph) ToAdjacencyList() [][]int {
	list := make([][]int, len(g.Vertices))

	for i, v := range g.Vertices {
		for _, n := range v.Neighbors {
			list[i] = append(list[i], n.Value)
		}
	}

	return list
}

// ToAdjacencyMatrix returns a graph represented as an Adjacency Matrix
func (g *Graph) ToAdjacencyMatrix() [][]int {
	matrix := make([][]int, len(g.Vertices))

	for i, _ := range g.Vertices {
		row := make([]int, len(g.Vertices))
		for j, _ := range g.Vertices {
			row[j] = 0
		}
		matrix[i] = row
	}

	for _, edge := range g.Edges() {
		matrix[edge[0]][edge[1]] = 1
	}

	return matrix
}
