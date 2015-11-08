package graph

import "testing"

func TestGraphCanHaveOneVertex(t *testing.T) {
	graph := NewGraph()
	vertex := graph.AddVertex(0)

	if vertex.Value != 0 {
		t.Errorf("Vertex value should be 0, found %d", vertex.Value)
	}

	if len(graph.Vertices) != 1 {
		t.Errorf("Graph length should be 1, found %d", len(graph.Vertices))
	}

	if graph.GetVertex(0) != vertex {
		t.Errorf("Did not get single vertex")
	}
}

func TestVertexOperations(t *testing.T) {
	graph := NewGraph()
	v1 := graph.AddVertex(0)
	v2 := graph.AddVertex(1)
	v3 := graph.AddVertex(2)
	v4 := graph.AddVertex(3)

	graph.AddEdge(v1, v2)
	graph.AddEdge(v2, v3)
	graph.AddEdge(v2, v4)

	// Test neighbors
	if len(v1.GetNeighbors()) != 1 {
		t.Errorf("First vertex should only have one neighbor")
	}

	if len(v2.GetNeighbors()) != 3 {
		t.Errorf("Second vertex should have three neighbors")
	}

	if len(v3.GetNeighbors()) != 1 {
		t.Errorf("Second vertex should only have one neighbor")
	}

	// Test Adjacency
	if v1.IsAdjacent(v2) != true {
		t.Errorf("Vertex 1 and vertex 2 should be adjacent")
	}

	if v1.IsAdjacent(v3) != false {
		t.Errorf("Vertex 1 and vertex 2 should be adjacent")
	}

	// Test can remove an edge
	v2.RemoveEdge(v1)

	if len(v1.GetNeighbors()) != 0 {
		t.Errorf("Vertex 1 should have no neighbors after removal")
	}

	if v2.IsAdjacent(v1) {
		t.Errorf("Vertex 2 should not be adjacent to vertex 1")
	}
}

func TestVertexOperationsOnDifferentGraphs(t *testing.T) {
	graph1 := NewGraph()
	graph2 := NewGraph()
	v1 := graph1.AddVertex(0)
	v2 := graph1.AddVertex(1)
	v3 := graph2.AddVertex(2)
	v4 := graph2.AddVertex(3)

	graph1.AddEdge(v1, v2)
	graph2.AddEdge(v3, v4)

	// Test Adjacency
	if v1.IsAdjacent(v3) != false {
		t.Errorf("Vertex 1 and vertex 3 are in different graphs and should not be adjacent")
	}

	// Test cannot remove an edge
	if v1.RemoveEdge(v3) != false {
		t.Errorf("Vertex 1 and vertex 3 are in different graphs and should not have an edge")
	}
}

func TestGraphCanFindAdjacentVertices(t *testing.T) {
	graph := NewGraph()
	v1 := graph.AddVertex(0)
	v2 := graph.AddVertex(1)

	graph.AddEdge(v1, v2)

	if graph.Adjacent(v1, v2) != true {
		t.Errorf("Vertex 1 and vertex 2 should be adjacent")
	}

	graph2 := NewGraph()
	v3 := graph2.AddVertex(0)
	v4 := graph2.AddVertex(1)

	graph2.AddEdge(v3, v4)

	if graph.Adjacent(v1, v3) != false {
		t.Errorf("Vertex 1 and vertex 2 are in different graphs and cannot be adjacent")
	}
}

func TestGraphCanGetEdgeList(t *testing.T) {
	graph := NewGraph()
	v1 := graph.AddVertex(0)
	v2 := graph.AddVertex(1)
	v3 := graph.AddVertex(2)
	v4 := graph.AddVertex(3)

	graph.AddEdge(v1, v2)
	graph.AddEdge(v2, v3)
	graph.AddEdge(v2, v4)

	edges := graph.Edges()

	if len(edges) != 6 {
		t.Errorf("There should be 6 edge pairs but found %d", len(edges))
	}
}

func TestGraphCanRemoveVertex(t *testing.T) {
	graph := NewGraph()
	v1 := graph.AddVertex(0)
	v2 := graph.AddVertex(1)
	v3 := graph.AddVertex(2)
	v4 := graph.AddVertex(3)

	graph.AddEdge(v1, v2)
	graph.AddEdge(v2, v3)
	graph.AddEdge(v2, v4)

	graph.RemoveVertex(v2)

	if len(graph.Vertices) != 3 {
		t.Errorf("Graph should have 3 vertices after removal, not %d", len(graph.Vertices))
	}

	if len(v1.GetNeighbors()) != 0 {
		t.Errorf("Vertex 1 should have no neighbors after removal, not %d", len(v1.GetNeighbors()))
	}

	if len(v3.GetNeighbors()) != 1 {
		t.Errorf("Vertex 3 should have 1 neighbor after removal, not %d", len(v3.GetNeighbors()))
	}

	if len(v4.GetNeighbors()) != 0 {
		t.Errorf("Vertex 4 should have no neighbors after removal, not %d", len(v4.GetNeighbors()))
	}

	if graph.GetVertex(1) != nil {
		t.Errorf("Should not be able to get removed vertex")
	}
}

func TestGraphCanManipulateEdges(t *testing.T) {
	graph := NewGraph()
	v1 := graph.AddVertex(0)
	v2 := graph.AddVertex(1)
	v3 := graph.AddVertex(2)
	v4 := graph.AddVertex(3)

	graph.AddEdge(v1, v2)
	graph.AddEdge(v2, v3)
	graph.AddEdge(v2, v4)

	graph.RemoveEdge(v1, v2)

	if len(v1.GetNeighbors()) != 0 {
		t.Errorf("Vertex 1 should have no neighbors after removal, not %d", len(v1.GetNeighbors()))
	}

	if len(v2.GetNeighbors()) != 2 {
		t.Errorf("Vertex 2 should have two neighbors after removal, not %d", len(v2.GetNeighbors()))
	}
}

func TestGraphCanBeRepresentation(t *testing.T) {
	graph := NewGraph()
	v1 := graph.AddVertex(0)
	v2 := graph.AddVertex(1)
	v3 := graph.AddVertex(2)
	v4 := graph.AddVertex(3)

	graph.AddEdge(v1, v2)
	graph.AddEdge(v2, v3)
	graph.AddEdge(v2, v4)

	// Test Adjacency List
	list := graph.ToAdjacencyList()
	targetList := [][]int{
		[]int{1},
		[]int{0, 2, 3},
		[]int{1},
		[]int{1},
	}

	if len(list) != 4 {
		t.Errorf("The adjacency list should have 4 vertices")
	}

	for i, row := range list {
		for j, item := range row {
			if item != targetList[i][j] {
				t.Errorf("Adjacency point at [%d, %d] should be %d but found $d", i, j, targetList[i][j], item)
			}
		}
	}

	// Test Adjacency Matrix
	matrix := graph.ToAdjacencyMatrix()
	targetMatrix := [][]int{
		[]int{0, 1, 0, 0},
		[]int{1, 0, 1, 1},
		[]int{0, 1, 0, 0},
		[]int{0, 1, 0, 0},
	}

	if len(matrix) != 4 {
		t.Errorf("The adjacency matrix should have 4 vertices")
	}

	for i, row := range matrix {
		for j, item := range row {
			if item != targetMatrix[i][j] {
				t.Errorf("Adjacency point at [%d, %d] should be %d but found $d", i, j, targetMatrix[i][j], item)
			}
		}
	}
}
