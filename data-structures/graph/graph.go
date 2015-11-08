// ----------------------------------------------------------------------------
// Data Structures - Graph
// ----------------------------------------------------------------------------
//
// A graph is a container which holds vertices or nodes. These vertices can be
// connected by different edges and when we have these edges, we can traverse
// a graph finding the shortest path to different vertices as well as creating
// lines between vertices. Graphs can be directed or undirected. The
// difference being that a directed graph has edges that go one way and
// undirected graphs can go both backwards and forwards along the edge.
//
// In the following code, I have implemented a undirected graph. As a handy
// way to view the graph and use it, I have provided two functions to return
// it:
//
// - ToAdjacencyList()
// - ToAdjacencyMatrix()
//
// An Adjacency list has vertices stored as records or objects, and every
// vertex stores a list of adjacent vertices. This data structure allows the
// storage of additional data on the vertices. Additional data can be stored
// if edges are also stored as objects, in which case each vertex stores its
// incident edges and each edge stores its incident vertices. An example is:
//
//		[1]
//		[0 2 3]
//		[1]
//		[1]
//
// An Adjacency Matrix is a two-dimensional matrix, in which the rows represent
// source vertices and columns represent destination vertices. Data on edges
// and vertices must be stored externally. Only the cost for one edge can be
// stored between each pair of vertices. An example is:
//
//
//		[0 1 0 0]
//		[1 0 1 1]
//		[0 1 0 0]
//		[0 1 0 0]
//
// Adjacency lists are generally preferred because they efficiently represent
// sparse graphs. An adjacency matrix is preferred if the graph is dense, that
// is the number of edges is close to the number of vertices squared or if one
// must be able to quickly look up if there is an edge connecting two vertices.
//
// ----------------------------------------------------------------------------
// Time Complexity
// ----------------------------------------------------------------------------
//
//					------------------------------------
//					| Adjacency List	Adjacency Matrix
//					------------------------------------
// Store graph		| O(|V|+|E|)		O(|V|^2)
// Add vertex		| O(1)				O(|V|^2)
// Add edge			| O(1)				O(1)
// Remove vertex	| O(|E|)			O(|V|^2)
// Remove edge		| O(|E|)			O(1)
//
// ----------------------------------------------------------------------------
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
	var removeIndex int

	if v.Graph != other.Graph {
		return false
	}

	if !v.IsAdjacent(other) {
		return false
	}

	for i, n := range v.Neighbors {
		if n == other {
			removeIndex = i
		}
	}

	if len(v.Neighbors) == 1 {
		v.Neighbors = []*Vertex{}
	} else {
		v.Neighbors = append(v.Neighbors[:removeIndex], v.Neighbors[removeIndex+1:]...)
	}

	for i, n := range other.Neighbors {
		if n == v {
			removeIndex = i
		}
	}

	if len(other.Neighbors) == 1 {
		other.Neighbors = []*Vertex{}
	} else {
		other.Neighbors = append(other.Neighbors[:removeIndex], other.Neighbors[removeIndex+1:]...)
	}

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
