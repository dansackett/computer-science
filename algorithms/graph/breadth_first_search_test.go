package graph

import "testing"

func TestBreadthFirstSearchWorks(t *testing.T) {
	graph := [][]int{
		[]int{1},
		[]int{0, 4, 5},
		[]int{3, 4, 5},
		[]int{2, 6},
		[]int{1, 2},
		[]int{1, 2, 6},
		[]int{3, 5},
		[]int{},
	}

	result := BreadthFirstSearch(graph, 3)

	if result[0]["distance"] != 4 && result[0]["parent"] != 1 {
		msg := "For vertex 0, distance should be 4 and parent should be 1. Found distance as %d and parent as %d"
		t.Errorf(msg, result[0]["distance"], result[0]["parent"])
	}

	if result[1]["distance"] != 3 && result[1]["parent"] != 4 {
		msg := "For vertex 1, distance should be 3 and parent should be 4. Found distance as %d and parent as %d"
		t.Errorf(msg, result[1]["distance"], result[1]["parent"])
	}

	if result[2]["distance"] != 1 && result[2]["parent"] != 3 {
		msg := "For vertex 2, distance should be 1 and parent should be 3. Found distance as %d and parent as %d"
		t.Errorf(msg, result[2]["distance"], result[2]["parent"])
	}

	if result[3]["distance"] != 0 && result[3]["parent"] != -1 {
		msg := "For vertex 3, distance should be 0 and parent should be -1. Found distance as %d and parent as %d"
		t.Errorf(msg, result[3]["distance"], result[3]["parent"])
	}

	if result[4]["distance"] != 2 && result[4]["parent"] != 2 {
		msg := "For vertex 4, distance should be 2 and parent should be 2. Found distance as %d and parent as %d"
		t.Errorf(msg, result[4]["distance"], result[4]["parent"])
	}

	if result[5]["distance"] != 2 && result[5]["parent"] != 2 {
		msg := "For vertex 5, distance should be 2 and parent should be 2. Found distance as %d and parent as %d"
		t.Errorf(msg, result[5]["distance"], result[5]["parent"])
	}

	if result[6]["distance"] != 1 && result[6]["parent"] != 3 {
		msg := "For vertex 6, distance should be 1 and parent should be 3. Found distance as %d and parent as %d"
		t.Errorf(msg, result[6]["distance"], result[6]["parent"])
	}

	if result[7]["distance"] != -1 && result[7]["parent"] != -1 {
		msg := "For vertex 7, distance should be -1 and parent should be -1. Found distance as %d and parent as %d"
		t.Errorf(msg, result[7]["distance"], result[7]["parent"])
	}
}
