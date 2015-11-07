// ----------------------------------------------------------------------------
// Algorithms - Breadth-First Search (Graph)
// ----------------------------------------------------------------------------
//
// Breadth-first search is a method of traversing and searching tree / graph
// data structures. The basic concept is starting at a base vertex, we check
// the nearest neighbors assigning a distance from the base vertex and the
// vertex which leads to it. We then visit the neighbors of those neighbors
// until each vertex has been visited.
//
// For example, if we have a graph represented as an adjacency list:
//
//     0: [1],
//     1: [0, 4, 5],
//     2: [3, 4, 5],
//     3: [2, 6],
//     4: [1, 2],
//     5: [1, 2, 6],
//     6: [3, 5],
//     7: []
//
// We can choose to start at vertex 3. Since it has no parent, we set it as a
// null-like value and set the distance to this as 0 since it's the base
// vertex. We then look to see the nearest neighbor vertices which in this
// case is vertex 2.
//
// With vertex 2, we set the distance to 1 and the parent to our base. Since
// there are no more neighbors at a distance of 1, we find the vertices which
// vertex 2 leds to. These are the new neighbors and they are vertex 4 and
// vertex 5.
//
// This pattern continues until each vertex has been assigned values. If we
// already visited the vertex, we skip it. In the end, you will have a mapping
// of each vertex with the distance from the base vertex and the vertex that
// leads to it. This is important because it allows us to:
//
// - Build a tree based on the distances
// - Find the shortest path from one vertex to another
// - Draw paths between vertices
//
// In this algorith, a Queue is used to keep track of the next vertices it
// needs to visit. This allows us to iterate through the neighbors and then
// pop those off the queue to get their next neighbors.
//
// ----------------------------------------------------------------------------
// Time Complexity
// ----------------------------------------------------------------------------
//
// In the worst case, the complexity is defined as O(|V| + |E|) where |V|
// represents the number of vertices and |E| represents the number of edges.
// With the adjacency list this expects, it can be defined as Θ(|V| + |E|).
//
// ----------------------------------------------------------------------------
package graph

import q "../../data-structures/queue"

func BreadthFirstSearch(graph [][]int, source int) []map[string]int {
	data := make([]map[string]int, len(graph))

	for i, _ := range graph {
		data[i] = map[string]int{
			"distance": -1,
			"parent":   -1,
		}
	}

	data[source]["distance"] = 0

	queue := q.NewQueue()
	queue.Enqueue(source)

	for !queue.IsEmpty() {
		currentVertex := queue.Dequeue()

		for i, _ := range graph[currentVertex] {
			neighbor := graph[currentVertex][i]

			if data[neighbor]["distance"] < 0 {
				data[neighbor]["distance"] = data[currentVertex]["distance"] + 1
				data[neighbor]["parent"] = currentVertex
				queue.Enqueue(neighbor)
			}
		}
	}

	return data
}
